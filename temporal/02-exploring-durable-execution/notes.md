# Temporal 102: Exploring Durable Execution

## Table of Contents

- [Welcome](#welcome)
  - [Prerequisites](#prerequisites)
  - [Course Goals](#course-goals)
  - [Example Code and Learning Approach](#example-code-and-learning-approach)
  - [Examples Used in the Course](#examples-used-in-the-course)
- [Course Outcomes](#course-outcomes)
  - [Evaluate Durable Execution](#evaluate-durable-execution)
  - [Apply Development Best Practices](#apply-development-best-practices)
  - [Debug Workflow Execution](#debug-workflow-execution)
  - [Prepare for Production](#prepare-for-production)
- [Durable Execution System](#durable-execution-system)
  - [What Is a Durable Execution System?](#what-is-a-durable-execution-system)
  - [Developer Productivity](#developer-productivity)
- [Temporal Application Structure](#temporal-application-structure)
  - [Workflows and Activities](#workflows-and-activities)
  - [Workers and Application Code](#workers-and-application-code)
  - [The Temporal Cluster](#the-temporal-cluster)
  - [Temporal Cloud](#temporal-cloud)
  - [Separation of Application and Cluster](#separation-of-application-and-cluster)
- [How Errors Affect Workflow Execution](#how-errors-affect-workflow-execution)
  - [Activity Errors](#activity-errors)
  - [Workflow Errors](#workflow-errors)
  - [Returning Errors from Application Code](#returning-errors-from-application-code)
  - [Cross-Language Error Handling](#cross-language-error-handling)
- [Backwards-Compatible Evolution of Input Parameters and Return Values](#backwards-compatible-evolution-of-input-parameters-and-return-values)
  - [Example: Struct-Based Activity Input and Output](#example-struct-based-activity-input-and-output)
  - [Compatibility Considerations](#compatibility-considerations)
  - [Optional Self-Study: Using Structs](#optional-self-study-using-structs)

## Welcome

### Prerequisites

This course is intended for experienced software developers with basic Go proficiency and foundational Temporal knowledge. **Temporal 101** is the recommended prerequisite.

You should already be able to:

- Develop and execute Workflows and Activities in Go.
- Configure and run a Worker.
- Navigate the Temporal Web UI.
- Explain the high-level interactions between a Temporal application and the Temporal Cluster during Workflow Execution.

### Course Goals

Temporal 102 moves beyond basic application development toward skills needed for production deployment. It focuses on:

- Understanding common problems encountered by Temporal developers and why they occur.
- Identifying, solving, and preventing those problems.
- Building a deeper understanding of Temporal through key concepts and best practices.

### Example Code and Learning Approach

The examples use more Go features than Temporal 101, including **structs, pointers, and interfaces**, as well as third-party packages such as **`testify`** for unit tests. Expert Go knowledge is not required.

The code is designed to teach specific Temporal concepts, rather than serve as a production-ready template. Some examples contain **intentional flaws** so that failures can be observed and understood in a learning environment.

### Examples Used in the Course

| Example | Learning purpose |
| :--- | :--- |
| Translation Workflow | Replaces Temporal 101's Spanish greeting/farewell microservice with support for other words and additional languages. The first exercise uses structs for business inputs and results to demonstrate best practices. |
| Pizza-order Workflow | Shows how Workflow code produces **Commands** sent to the Temporal Cluster, which lead to **History Events**. These concepts underpin history replay and recovery after a Worker crash. |
| Loan-processing Workflow | Demonstrates a long-running Workflow, why Workflow code must be **deterministic**, and what can happen when it is not. |

## Course Outcomes

The course covers Temporal across the **full development lifecycle**: testing, debugging, deploying, and updating applications. Its objectives fall into four areas.

### Evaluate Durable Execution

- Explain how **History Replay** recreates the state of a Workflow Execution.
- Use **Timers** to introduce delays and determine how a Worker crash affects them.
- Summarize the relationship between **Temporal SDK API calls, Worker Commands, and History Events**.
- Explain determinism in the context of Workflow Execution and its implications for application code.

### Apply Development Best Practices

- Model business input parameters and return values as **data structures** to support backwards-compatible evolution.
- Use logging to report information from Workflows and Activities during execution.
- Identify, correct, and avoid common sources of Workflow non-determinism.
- Select Workflow IDs appropriate to the use case.
- Develop and execute unit tests for Workflows and Activities.
- Use **mocks** to isolate Workflow tests from Activity implementations.

### Debug Workflow Execution

- Analyze **Event History** to diagnose execution problems.
- Use the Web UI to trace both open and closed Workflow Executions.
- View the stack trace of a current Workflow Execution.
- Distinguish the different states of Workflows and Activities.
- Explain the role of **Sticky Task Queues** in Workflow Execution.
- Describe how returning errors from Workflows and Activities affects execution.
- Observe the consequences of running non-deterministic Workflow code.
- **Reset** a Workflow Execution to recover from a bad deployment.
- Distinguish **cancellation** from **termination**.
- Terminate Workflows using the CLI, SDK, and Web UI.

### Prepare for Production

- Illustrate the logical and physical views of a production Temporal application deployment.
- Configure a Temporal Client to connect to **Temporal Cloud** or a secure self-hosted cluster.
- Identify which changes can be safely deployed **without versioning**.

## Durable Execution System

### What Is a Durable Execution System?

Temporal is a **durable execution system**: it preserves application execution state so code can recover automatically from failures and continue making progress.

Failures can range from a **network timeout** to a **kernel panic on an application server**. Rather than requiring the application to reconstruct its progress from scratch, Temporal retains the information needed to recover Workflow state.

Durable execution supports reliable orchestration; application code still needs correct business logic and retry-safe external side effects.

### Developer Productivity

Developers normally spend substantial effort handling failures and timeouts. Temporal provides **higher-level abstractions**, recovery mechanisms, and built-in scalability so more of that effort can go toward business logic.

The **Web UI** also makes execution observable. It lets you inspect past and current Workflow Executions, including their:

- Event History.
- Input parameters.
- Return values, when available.

In this course, the Web UI is used to **debug Workflow Execution and verify that a fix resolves the problem**.

## Temporal Application Structure

### Workflows and Activities

A **Workflow** defines the sequence of steps that makes up an application's main business logic. Temporal provides language-specific **SDKs** with APIs and libraries for implementing applications in general-purpose languages. In this course, Workflow and Activity Definitions are Go functions using the Go SDK.

Workflow code must be **deterministic**. “The same input produces the same output” is an introductory explanation, but the more precise requirement is **replay compatibility: Workflow code must produce the sequence of Commands expected by the recorded Event History**. The course explores why this matters, what happens when it is violated, and how to avoid non-determinism.

**Activities** encapsulate business logic that is non-deterministic or prone to failure, such as external I/O. Workflows schedule Activities, and Temporal automatically retries retryable failures according to their Retry Policy and applicable timeouts. This handles transient failures without a custom retry loop in Workflow code.

### Workers and Application Code

The **Temporal Cluster orchestrates execution; it does not execute application code**. Workers poll Task Queues and execute Workflow and Activity code.

The SDK provides the Worker implementation. The developer writes code to:

1. Configure a Worker with a Task Queue name.
2. Register the Workflow and Activity Definitions it can execute.
3. Start the Worker.

The course uses **Temporal Application Code** to collectively mean:

- Workflow Definitions.
- Activity Definitions.
- Worker configuration and startup code.

A complete Temporal application includes both this developer-written code and the supporting functionality supplied by the SDK.

### The Temporal Cluster

A self-hosted **Temporal Cluster** consists of the Temporal Server—its Frontend Service and backend services—plus the database used for persistence. Its support enables durable, scalable, and reliable Workflow Execution.

Optional supporting components can include:

- **Elasticsearch** for advanced visibility, depending on the deployment's visibility configuration; supported SQL visibility is another option.
- **Grafana** for operational dashboards showing cluster and application health, backed by a metrics data source.

### Temporal Cloud

**Temporal Cloud** provides a managed service fulfilling the same orchestration and persistence roles as a self-hosted cluster. It removes the need for your team to deploy, secure, operate, and maintain the Temporal Server infrastructure.

The operational responsibilities differ, but the core Workflow programming and execution model is the same. In this course, **“Temporal Cluster” generally also applies to Temporal Cloud**, unless a distinction is explicitly made.

Moving an application from a self-hosted cluster to Temporal Cloud typically requires minimal code changes, primarily to the Temporal Client's connection settings, such as endpoint, namespace, and authentication. The course covers these settings later.

### Separation of Application and Cluster

**Application code executes outside the Temporal Cluster**, whether the cluster is self-hosted or provided by Temporal Cloud.

```text
Temporal application                         Temporal Cluster / Cloud
┌──────────────────────────────┐             ┌──────────────────────────┐
│ Developer-written code       │             │ Frontend Service         │
│ • Workflow Definitions       │  polls and  │ Backend services         │
│ • Activity Definitions       │◄───────────►│ Task Queues              │
│ • Worker configuration       │ reports     │ Persistent Event History │
│                              │ results     │ and execution state      │
│ SDK-provided Worker + Client │             │                          │
└──────────────────────────────┘             └──────────────────────────┘
```

In production, Workers and the cluster typically run on separate machines or containers and can even run in different data centers. Workers must be able to reach the cluster's Frontend Service. Using Temporal Cloud manages the cluster infrastructure; your team still deploys and operates the application Workers.

## How Errors Affect Workflow Execution

Returning an error signals failure, but **Activity attempt failure and Workflow Execution failure have different consequences**.

| Where the error is returned | Default behavior |
| :--- | :--- |
| Activity Definition | The attempt fails; Temporal retries retryable errors according to the Activity's Retry Policy and applicable timeouts. |
| Workflow Definition | The Workflow Execution fails; there is no Workflow Retry Policy by default. |

### Activity Errors

Activities perform error-prone operations, such as querying a database. Failures might involve a temporary outage, rejected credentials, or invalid SQL.

The default Activity Retry Policy uses **exponential backoff**:

```text
1s → 2s → 4s → 8s → 16s → 32s → 64s → 100s → 100s → …
```

These are delays between attempts, not the duration of the attempts themselves. By default, there is no attempt-count limit. Retries can stop when the Activity succeeds, is canceled, encounters a non-retryable error, reaches an applicable overall timeout, or its Workflow closes. A Start-to-Close Timeout limits one attempt and can trigger another retry; it does not limit the entire retry sequence.

Temporary outages may resolve without intervention. Persistent problems require a fix: for example, correct credentials or SQL, then deploy and restart the Activity Workers if code changed. A subsequent attempt can use the corrected implementation, provided retries are still allowed and the updated Worker receives the task.

Customize retries through **`workflow.ActivityOptions.RetryPolicy`**. Different calls to `workflow.ExecuteActivity` can use different options and policies to match their business requirements.

**One `ExecuteActivity` call can result in multiple executions of the Activity code.** Make external side effects **idempotent**: repeating the same logical operation should have the same effect as performing it once. For operations such as payments, this commonly requires an idempotency key or deduplication in the external system.

While retries continue, the Activity Future remains pending. If the Activity ultimately fails, its Future's `Get` returns an error to Workflow code. The Workflow can handle that error with a fallback or compensation; Activity failure does not automatically mean the Workflow must fail.

### Workflow Errors

Workflows have **no Retry Policy by default**, and adding one is unusual. If the Workflow Definition returns an error, the Workflow Execution is marked as failed and is not automatically retried under the default configuration.

Activities are retried because external operations often encounter transient failures. Workflow code instead performs deterministic orchestration and should not directly perform those external operations. It decides how to handle an Activity's eventual failure and whether to return an error itself.

Prefer correcting a recoverable underlying problem while execution can still make progress, rather than failing the Workflow unnecessarily. Return an error when failure is the intended business outcome or recovery cannot be completed.

**Workflow Execution failure is distinct from Workflow Task failure.** A Worker crash or a Workflow-code problem such as a non-determinism error does not, by itself, mean the Workflow Definition returned an error and closed the execution as failed. Such problems can prevent progress while the execution remains open.

### Returning Errors from Application Code

Use ordinary Go errors; a Temporal-specific API is not required. For example, an Activity making an HTTP request can return the original error:

```go
resp, err := http.Get(url)
if err != nil {
	return "", err
}
defer resp.Body.Close()
```

This fragment assumes a function returning `(string, error)` and continues with response handling after the shown code. Returning the original error preserves its diagnostic information, whereas replacing it with `errors.New("request failed")` discards the cause. In an Activity, prefer an HTTP request using the Activity's context so cancellation propagates to the request.

The SDK's **`temporal` package** also provides error constructors for explicit application error types and **non-retryable errors**. Use these when the failure needs semantics beyond a generic Go error, such as indicating that another Activity attempt cannot resolve invalid input.

### Cross-Language Error Handling

The Worker automatically converts returned errors into Temporal's **language-neutral failure representation**. The cluster can persist failure details consistently, and SDKs can decode them into their own language's error model.

This supports applications spanning multiple SDKs—for example, a TypeScript client starting a Go Workflow that schedules Python and Java Activities. Each Activity still needs a suitable Worker and Task Queue routing; language-neutral failures allow errors to cross these boundaries without depending on a shared native error type.

## Backwards-Compatible Evolution of Input Parameters and Return Values

**Use a single struct for business input and a struct for the result** of Workflow and Activity Definitions. This lets their data evolve without changing the function signature.

- A Workflow must take **`workflow.Context`** as its first parameter.
- An Activity should take **`context.Context`** as its first parameter, although it is not required. It enables cancellation notification and other Activity SDK features.
- These SDK-supplied contexts are separate from the serialized business input.

Both definitions can accept multiple business parameters, but changing their **number, order, or types** can break compatibility with existing executions and callers. Adding fields to an input struct avoids changing the parameter list. A result struct similarly allows additional output fields without changing the return type.

### Example: Struct-Based Activity Input and Output

Temporal 101's greeting Activity accepted a name and returned a Spanish greeting:

```go
func GreetInSpanish(ctx context.Context, name string) (string, error)
```

Supporting additional languages by adding a separate `languageCode string` parameter would change that signature. Starting with structs gives the data room to evolve:

```go
type GreetingInput struct {
	Name         string
	LanguageCode string
}

type GreetingOutput struct {
	Greeting string
}
```

The Activity now has a stable struct-based signature. With `context` and `fmt` imported, this illustrative implementation supports French and returns an error for other language codes:

```go
func GetTranslatedGreeting(ctx context.Context, input GreetingInput) (GreetingOutput, error) {
	if input.LanguageCode == "fr" {
		return GreetingOutput{
			Greeting: fmt.Sprintf("Bonjour, %s", input.Name),
		}, nil
	}

	return GreetingOutput{}, fmt.Errorf("unsupported language code: %q", input.LanguageCode)
}
```

Additional languages can be implemented without adding function parameters. Additional result fields can be added to `GreetingOutput` as requirements evolve. In a real Activity, classify unsupported input appropriately if retrying it cannot help; the ordinary error above is retryable by default.

### Compatibility Considerations

**Structs support compatible evolution; they do not make every change compatible.** With the default JSON encoding, adding a field is generally compatible at the serialization level, but the implementation must still handle older payloads:

- A field absent from stored input decodes to its Go **zero value**. Define a suitable default or otherwise preserve the behavior expected by older executions.
- Renaming fields, changing their types, or making a new field mandatory can break compatibility.
- Result changes must also remain compatible with callers and Workflows consuming older recorded results.
- Keep serialized fields exported so the default JSON converter can encode them.

The initial migration from `string` inputs/results to structs **is not backwards-compatible** with the old payload shape. Adopt the struct-based design early rather than assuming an existing string-based execution can decode into the new type.

Payload compatibility is also separate from Workflow determinism: a stable function signature does not make a change to the Workflow's Command sequence replay-compatible.

### Optional Self-Study: Using Structs

The local **`samples/using-structs/`** example demonstrates struct-based inputs and results for both Workflow and Activity Definitions. It contains `microservice/`, `start/`, and `worker/` directories for exploring the complete application flow.
