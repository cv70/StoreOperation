import { useEffect, useState, type ReactNode } from 'react'
import './App.css'

type Role = 'store_manager' | 'supervisor' | 'hq'

type Metric = {
  key: string
  label: string
  value: string
  trend: string
}

type Card = {
  id: string
  title: string
  reason: string
  action: string
  impact: string
  deadline: string
  priority: string
  status: string
}

type Overview = {
  role: Role
  headline: string
  metrics: Metric[]
  opportunities: Card[]
  risks: Card[]
  focus: string[]
}

type ApiResponse<T> = {
  code: number
  msg?: string
  data: T
}

type ActionEvent = {
  card_id: string
  from_state: string
  to_state: string
  reason: string
  created_at?: string
}

function App() {
  const [role, setRole] = useState<Role>('store_manager')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')
  const [overview, setOverview] = useState<Overview | null>(null)
  const [updatingId, setUpdatingId] = useState('')
  const [historyByCard, setHistoryByCard] = useState<Record<string, ActionEvent[]>>({})
  const [historyCardId, setHistoryCardId] = useState('')
  const [historyLoadingId, setHistoryLoadingId] = useState('')
  const [historyPageByCard, setHistoryPageByCard] = useState<Record<string, number>>({})

  useEffect(() => {
    let cancelled = false

    const load = async () => {
      setLoading(true)
      setError('')
      try {
        const resp = await fetch(`/api/v1/workbench/overview?role=${role}`)
        const json = (await resp.json()) as ApiResponse<Overview>
        if (!resp.ok || json.code !== 200) {
          throw new Error(json.msg || '加载失败')
        }
        if (!cancelled) {
          setOverview(json.data)
        }
      } catch (err) {
        if (!cancelled) {
          setError(err instanceof Error ? err.message : '加载失败')
        }
      } finally {
        if (!cancelled) {
          setLoading(false)
        }
      }
    }

    load()

    return () => {
      cancelled = true
    }
  }, [role])

  const transition = async (cardId: string, toState: string, reason = '') => {
    setUpdatingId(cardId)
    try {
      const resp = await fetch(`/api/v1/workbench/actions/${cardId}/transition`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ to_state: toState, reason }),
      })
      const json = (await resp.json()) as ApiResponse<Card>
      if (!resp.ok || json.code !== 200) {
        throw new Error(json.msg || '更新失败')
      }
      setOverview((prev) => {
        if (!prev) {
          return prev
        }
        const patch = (cards: Card[]) => cards.map((c) => (c.id === cardId ? { ...c, status: json.data.status } : c))
        return {
          ...prev,
          opportunities: patch(prev.opportunities),
          risks: patch(prev.risks),
        }
      })
    } catch (err) {
      setError(err instanceof Error ? err.message : '更新失败')
    } finally {
      setUpdatingId('')
    }
  }

  const actionButtons = (card: Card) => {
    const disabled = updatingId === card.id
    const controls: ReactNode[] = []

    if (card.status === 'new') {
      controls.push(<button key="accept" disabled={disabled} onClick={() => transition(card.id, 'accepted')}>采纳</button>)
      controls.push(<button key="escalate" disabled={disabled} onClick={() => transition(card.id, 'escalated')}>升级</button>)
    }
    if (card.status === 'accepted') {
      controls.push(<button key="start" disabled={disabled} onClick={() => transition(card.id, 'in_progress')}>开始执行</button>)
    }
    if (card.status === 'in_progress') {
      controls.push(
        <button key="done" disabled={disabled} onClick={() => transition(card.id, 'done_pending_validation')}>
          完成待验证
        </button>,
      )
    }

    controls.push(
      <button
        key="history"
        className="secondary"
        disabled={historyLoadingId === card.id}
        onClick={() => toggleHistory(card.id)}
      >
        {historyCardId === card.id ? '收起历史' : '查看历史'}
      </button>,
    )

        return (
      <>
        <div className="buttons">{controls}</div>
        {!['new', 'accepted', 'in_progress'].includes(card.status) && <span className="status-tag">{card.status}</span>}
      </>
    )
  }

  const toggleHistory = async (cardId: string) => {
    if (historyCardId === cardId) {
      setHistoryCardId('')
      return
    }
    if (!historyByCard[cardId]) {
      setHistoryLoadingId(cardId)
      try {
        const resp = await fetch(`/api/v1/workbench/actions/${cardId}/events`)
        const json = (await resp.json()) as ApiResponse<ActionEvent[]>
        if (!resp.ok || json.code !== 200) {
          throw new Error(json.msg || '历史加载失败')
        }
        setHistoryByCard((prev) => ({ ...prev, [cardId]: json.data }))
      } catch (err) {
        setError(err instanceof Error ? err.message : '历史加载失败')
        setHistoryLoadingId('')
        return
      } finally {
        setHistoryLoadingId('')
      }
    }
    setHistoryPageByCard((prev) => ({ ...prev, [cardId]: prev[cardId] || 1 }))
    setHistoryCardId(cardId)
  }

  const historyBlock = (cardId: string) => {
    const events = historyByCard[cardId] || []
    const page = historyPageByCard[cardId] || 1
    const pageSize = 5
    const start = (page - 1) * pageSize
    const pageEvents = events.slice(start, start + pageSize)
    const totalPages = Math.max(1, Math.ceil(events.length / pageSize))
    if (events.length === 0) {
      return <p className="history-empty">暂无历史记录</p>
    }
    return (
      <>
        <ul className="history-list">
          {pageEvents.map((event, idx) => (
            <li key={`${event.card_id}-${event.to_state}-${idx}`}>
              <span>{event.from_state} → {event.to_state}</span>
              {event.reason && <span> · {event.reason}</span>}
              {event.created_at && <span> · {event.created_at}</span>}
            </li>
          ))}
        </ul>
        <div className="history-pager">
          <button
            className="secondary"
            disabled={page <= 1}
            onClick={() => setHistoryPageByCard((prev) => ({ ...prev, [cardId]: (prev[cardId] || 1) - 1 }))}
          >
            上一页
          </button>
          <span>{page} / {totalPages}</span>
          <button
            className="secondary"
            disabled={page >= totalPages}
            onClick={() => setHistoryPageByCard((prev) => ({ ...prev, [cardId]: (prev[cardId] || 1) + 1 }))}
          >
            下一页
          </button>
        </div>
      </>
    )
  }

  return (
    <main className="page">
      <header className="hero">
        <div>
          <p className="eyebrow">StoreOperation</p>
          <h1>AI + 软件 + 硬件门店工作台</h1>
          <p className="subtitle">先执行，再复盘，再放大。</p>
        </div>
        <div className="role-switch">
          <button className={role === 'store_manager' ? 'active' : ''} onClick={() => setRole('store_manager')}>
            店长台
          </button>
          <button className={role === 'supervisor' ? 'active' : ''} onClick={() => setRole('supervisor')}>
            督导台
          </button>
          <button className={role === 'hq' ? 'active' : ''} onClick={() => setRole('hq')}>
            总部台
          </button>
        </div>
      </header>

      {loading && <p className="state">加载中...</p>}
      {error && <p className="state error">{error}</p>}

      {overview && !loading && !error && (
        <div className="board-shell">
          <section className="board">
            <p className="headline">{overview.headline}</p>
            <div className="metrics">
              {overview.metrics.map((metric) => (
                <article key={metric.key} className="metric-card">
                  <p>{metric.label}</p>
                  <strong>{metric.value}</strong>
                  <span>{metric.trend}</span>
                </article>
              ))}
            </div>

            <div className="content-grid">
              <section className="panel">
                <h2>利润机会</h2>
                {overview.opportunities.map((card) => (
                  <article key={card.id} className="action-card">
                    <h3>{card.title}</h3>
                    <p>{card.reason}</p>
                    <p className="meta">{card.action}</p>
                    <p className="meta">
                      {card.impact} · {card.deadline}
                    </p>
                    <p className="meta">状态：{card.status}</p>
                    {actionButtons(card)}
                  </article>
                ))}
              </section>

              <section className="panel">
                <h2>高危风险</h2>
                {overview.risks.map((card) => (
                  <article key={card.id} className="action-card risk">
                    <h3>{card.title}</h3>
                    <p>{card.reason}</p>
                    <p className="meta">{card.action}</p>
                    <p className="meta">
                      {card.impact} · {card.deadline}
                    </p>
                    <p className="meta">状态：{card.status}</p>
                    {actionButtons(card)}
                  </article>
                ))}
              </section>
            </div>

            <section className="panel">
              <h2>本角色执行重点</h2>
              <ul>
                {overview.focus.map((item) => (
                  <li key={item}>{item}</li>
                ))}
              </ul>
            </section>
          </section>

          {historyCardId && (
            <aside className="history-drawer">
              <div className="history-head">
                <h3>动作历史</h3>
                <button className="secondary" onClick={() => setHistoryCardId('')}>关闭</button>
              </div>
              <p className="meta">卡片ID：{historyCardId}</p>
              {historyBlock(historyCardId)}
            </aside>
          )}
        </div>
      )}
    </main>
  )
}

export default App
