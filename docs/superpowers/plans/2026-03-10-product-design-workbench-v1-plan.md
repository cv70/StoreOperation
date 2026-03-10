# Product Design Workbench V1 Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Turn the approved product design direction into the next set of concrete product design artifacts, starting with the shared action card component and role-based homepage flows.

**Architecture:** Treat the design spec as the source of truth. Refine one core object first, then derive screen-level behavior from it. Avoid parallel redesign of unrelated modules until the action card, alert logic, and homepage priorities are stable.

**Tech Stack:** Markdown, existing PRD and product strategy docs

---

## Chunk 1: Action Card First

### Task 1: Define the action card component spec

**Files:**
- Reference: `docs/superpowers/specs/2026-03-10-product-design-workbench-v1.md`
- Create: `docs/superpowers/specs/2026-03-10-action-card-component-spec.md`

- [ ] **Step 1: Review the approved workbench design spec**

Read: `docs/superpowers/specs/2026-03-10-product-design-workbench-v1.md`
Expected: Clear understanding of the six required action card fields and role-specific usage.

- [ ] **Step 2: Write the action card component spec**

Include:
- field definitions
- state transitions
- role-specific variants
- alert vs opportunity differences
- AI output structure

- [ ] **Step 3: Review the spec for consistency**

Run: `sed -n '1,260p' docs/superpowers/specs/2026-03-10-action-card-component-spec.md`
Expected: A design artifact that can drive both product and implementation decisions.

## Chunk 2: Homepage Flows

### Task 2: Define the three homepage flows from the action card

**Files:**
- Reference: `docs/superpowers/specs/2026-03-10-product-design-workbench-v1.md`
- Create: `docs/superpowers/specs/2026-03-10-homepage-flow-spec.md`

- [ ] **Step 1: Map homepage priorities for each role**

Write exact first-screen priorities for:
- store manager
- supervisor
- HQ

- [ ] **Step 2: Define first-click and completion paths**

Document:
- what the user sees first
- first click path
- success path
- interruption path

- [ ] **Step 3: Review for redundancy and scope control**

Expected: No homepage becomes a generic dashboard or duplicate of another role.
