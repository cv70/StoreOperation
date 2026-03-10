# Upgrade Customer Sales One-Pager Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Rewrite the customer-facing sales one-pager so it reflects the approved AI + software + hardware positioning and sharper buyer-specific value story.

**Architecture:** Keep the work scoped to documentation. Use the approved positioning design as the source of truth, then rewrite the existing one-pager around a tighter headline, clearer moat narrative, and more concrete sales structure for CEO and HQ operations buyers.

**Tech Stack:** Markdown, existing product and business docs

---

## Chunk 1: Source Alignment

### Task 1: Lock the source documents and rewrite target

**Files:**
- Modify: `docs/sales-onepager-p0.md`
- Reference: `docs/superpowers/specs/2026-03-10-customer-sales-positioning-design.md`
- Reference: `README.md`

- [ ] **Step 1: Review the approved positioning spec**

Read: `docs/superpowers/specs/2026-03-10-customer-sales-positioning-design.md`
Expected: Clear guidance for product definition, moat, message hierarchy, and one-pager requirements.

- [ ] **Step 2: Review the current one-pager and main README**

Run: `sed -n '1,240p' docs/sales-onepager-p0.md`
Run: `sed -n '1,220p' README.md`
Expected: Current messaging and product baseline are understood before rewriting.

- [ ] **Step 3: Rewrite the one-pager around the new message structure**

Update the document to include:
- stronger headline centered on profit recovery
- elevator pitch explaining AI + software + hardware integration
- value pillars ordered as hardware visibility -> AI decisions -> software execution -> ROI proof
- explicit “not ordinary SaaS” section
- sharper buyer fit and objection handling

- [ ] **Step 4: Review the updated copy for consistency**

Run: `sed -n '1,260p' docs/sales-onepager-p0.md`
Expected: Messaging matches the approved spec and does not contradict existing product docs.

- [ ] **Step 5: Commit**

```bash
git add docs/superpowers/specs/2026-03-10-customer-sales-positioning-design.md docs/superpowers/plans/2026-03-10-upgrade-customer-sales-onepager.md docs/sales-onepager-p0.md
git commit -m "docs: sharpen customer sales positioning"
```
