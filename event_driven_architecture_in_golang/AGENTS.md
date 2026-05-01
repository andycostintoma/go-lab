# AGENTS.md

## Purpose

This repository is a study workspace for the book *Event-Driven Architecture in Golang*.

The goal is not to treat `code/` as one continuously evolving production app. It is a chapter-by-chapter learning repo made of separate snapshots, plus two note files:

- `Event_Driven_Architecture_In_Golang.md`: long-form book notes
- `Summary.md`: condensed chapter summaries aligned to the book structure
- `code/`: practical chapter snapshots of the MallBots system

Future sessions should preserve that framing.

## How To Work In This Repo

When answering questions or making updates:

1. Read the relevant section in `Summary.md`.
2. Read the matching section in `Event_Driven_Architecture_In_Golang.md`.
3. Read the corresponding chapter folder in `code/`.
4. Synthesize the idea by tying book concept -> summary -> concrete code.
5. Ask coaching questions grounded in all three sources.
6. Nudge the user toward the answer instead of immediately giving it away.
7. Only update `Summary.md` when it is missing, unclear, or can be materially improved.

Do not explain a chapter from code alone when the user is clearly following the book.
Do not treat the study flow as lecture-first by default; prefer synthesis-first, then coached questions.

## Important Repo Shape

- `code/` contains self-contained chapter folders, each with its own `go.mod`.
- These folders are snapshots, not one app that should be refactored globally.
- Folder numbering is not always a perfect 1:1 match with book chapter numbers.

## Known Chapter/Folder Mapping Notes

- `code/01_Designing_And_Planning` aligns with the early planning/design material.
- `code/02_Modular_Monolith` and `code/03_Domain_Events` both sit under the broader modular monolith/domain-events progression from the book.
- `code/06_Event_Carried_State_Transfer` corresponds to the event-carried state transfer chapter.
- `code/07_Message_Workflows_Orchestrated_Saga` is the renamed folder for the message workflows / orchestrated saga chapter.
- `code/08_Transactional_Messaging` is the local repo snapshot that corresponds to the book's Chapter 9, Transactional Messaging.

Use the actual local folder names when citing repo paths.

## Key Concepts Already Established

Keep these distinctions clear in future explanations:

- domain events: internal in-process coordination inside a bounded context
- event-sourcing events: persisted aggregate history
- integration events/messages: cross-module communication contracts

Also preserve this chapter distinction:

- event-carried state transfer is about sharing enough data for downstream local models
- message workflows / orchestrated saga is about coordinating long-running multi-step business work with commands, replies, and compensation
- transactional messaging is about reliability of local writes plus message publication/consumption, especially dual writes, inbox, and outbox

## Current Naming Decisions

- The folder `code/07_Message_Workflows` was intentionally renamed to `code/07_Message_Workflows_Orchestrated_Saga`.
- References in `README.md` and `Summary.md` were updated to use that new name.
- Do not rename folders again unless explicitly asked.

## Summary.md Expectations

`Summary.md` should stay aligned with the book.

That means:

- keep book chapter titles
- keep book section titles when the user asks for chapter summaries in the same style
- include book images/code blocks when asked to mirror the chapter structure closely
- keep the closing sections consistent across chapters: `## Key Takeaways`, `## Repo Anchors`, `## Further Reading`

`Repo Anchors` should point to the actual local snapshot folders, even when local folder names differ from the book chapter numbering.

## Recommended Reading Flow For Explanations

If a future session needs to explain a chapter, prefer this order:

1. state the chapter's main problem
2. state what new mechanism the chapter adds
3. show what changed from the previous chapter
4. point to the small set of files that demonstrate the change
5. ask a few guided questions to test the understanding
6. only then explain supporting infrastructure where needed

This repo is easiest to understand as a sequence of architectural changes, not as isolated files.

## Chapter 8 Mental Model

For `code/07_Message_Workflows_Orchestrated_Saga`, the minimum understanding needed is:

1. `ordering` no longer coordinates the full order-creation flow directly.
2. `ordering` creates the order locally and publishes `OrderCreated`.
3. `cosec` hears `OrderCreated` and starts the saga.
4. the saga sends commands to other modules and waits for replies.
5. failures trigger compensation such as `RejectOrder` and `CancelShoppingList`.

## Chapter 9 Mental Model

For `code/08_Transactional_Messaging`, the minimum understanding needed is:

1. sagas solve workflow reliability, not local dual-write reliability.
2. Chapter 9 focuses on transactional boundaries.
3. one request/message handling flow should share one DB transaction.
4. inbox handles idempotent consumption.
5. outbox persists outgoing messages atomically with local state changes.
6. an outbox processor publishes later from durable storage.

## Editing Guidance

- Prefer small, local edits.
- Do not assume all chapter folders should be made consistent with each other; they intentionally show architectural evolution.
- When updating summaries, verify against both `Summary.md` and `Event_Driven_Architecture_In_Golang.md`.
- When citing code, prefer the smallest set of files that shows the idea clearly.

## Good Default Starting Point

If a new session starts without much context, begin by asking or determining:

- which book chapter the user is on
- whether they want the book explanation, the code explanation, or both
- which local folder matches that chapter

Then work from `Summary.md` -> book chapter -> code snapshot.

## Current Progress

- The user has already worked through the distinction between `code/06_Event_Carried_State_Transfer` and `code/07_Message_Workflows_Orchestrated_Saga`.
- The main Chapter 8 takeaway already established is: order creation changed from direct cross-module calls to an orchestrated saga driven by `cosec`, commands, replies, and compensation.
- `Summary.md` was updated for Chapter 9 and aligned to the book structure, including the closing sections `## Key Takeaways`, `## Repo Anchors`, and `## Further Reading`.
- The local repo folder for book Chapter 9 is `code/08_Transactional_Messaging`, not `code/Chapter09`.

## Where We Left Off

- The user is currently on Chapter 9, especially the transactional-boundary / DI-container part.
- The latest concept explained was why the chapter introduces a scoped DI container so one request/message flow can share one DB transaction.
- The next likely question area is comparing the book's scoped-container approach with more idiomatic Go alternatives such as explicit unit-of-work / `WithinTx(...)` patterns.
- If the next session resumes explanation mode, start from Chapter 9 in this order: `Summary.md` chapter section -> matching section in `Event_Driven_Architecture_In_Golang.md` -> `code/08_Transactional_Messaging` files.
