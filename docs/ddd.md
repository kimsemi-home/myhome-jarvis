# DDD Boundaries

myhome-jarvis uses DDD vocabulary to make semantic ownership explicit without
adding framework-heavy layers. The executable SSOT owns the bounded contexts,
canonical concepts, aliases, owners, generated targets, and related concepts.

Bounded contexts:

- HomeControl: local commands, dry-run execution plans, daemon and Flutter command surfaces.
- HouseholdFinance: fixture-first household finance entities and transaction IR.
- CommerceIntelligence: fixture-first purchase IR, recurring candidates, and commerce recommendations.
- StorageLake: local lake paths, retention, and generated storage policy.
- SecurityPolicy: public-repo safety, auth boundaries, and allowed-language policy.
- AgentOps: Linear/offline work queue, planner, loop, scheduler, quality, and checkpoints.
- KnowledgeIndex: local lexical concept index and semantic duplication checks.

The project uses DDD terms as vocabulary, not as mandatory class names:
Entity, ValueObject, Aggregate, DomainEvent, Repository, Policy, Port, Adapter,
and AntiCorruptionLayer. New semantic terms should be registered in
the split `lisp/ssot/ddd*.lisp` registry before implementation and regenerated
into `generated/concepts.generated.json`.

Each canonical concept declares a `ddd_kind`, and the registry must use every
approved DDD kind at least once. The same executable SSOT also defines domain
events and harness case contracts, so generated concept evidence covers
semantic ownership, emitted events, and the fixture/golden checks that protect
the context.

`mhj ddd verify` checks duplicate concepts, alias drift, invalid bounded
contexts, invalid DDD kinds, missing or invalid domain events, missing harness
case contracts, missing generated targets, and KnowledgeIndex policy
constraints.
