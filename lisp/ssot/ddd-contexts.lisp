(in-package #:myhome-jarvis.ssot)

(defparameter *bounded-contexts*
  #((:name "HomeControl"
     :owner "commands"
     :description "Local home-control commands, dry-run execution plans, and Flutter command surface.")
    (:name "HouseholdFinance"
     :owner "finance"
     :description "Fixture-first household finance entities, transaction IR, and recommendation inputs.")
    (:name "CommerceIntelligence"
     :owner "commerce"
     :description "Fixture-first commerce purchase IR, recurring purchase candidates, and price intelligence.")
    (:name "ConnectorReadiness"
     :owner "connectors"
     :description "Public-safe planned connector catalog and read-only readiness status.")
    (:name "AgentCluster"
     :owner "agent-cluster"
     :description "Evidence-first multi-agent learning loop policy, authority gates, and verification sidecars.")
    (:name "StorageLake"
     :owner "storage"
     :description "Local lake layout, retention policy, and generated storage contracts.")
    (:name "SecurityPolicy"
     :owner "security"
     :description "Public-repo safety, secret scanning, auth boundaries, and allowed-language policy.")
    (:name "AgentOps"
     :owner "planner"
     :description "Closed-loop planner, Linear/offline work queue, scheduler, quality, and checkpoint evidence.")
    (:name "KnowledgeIndex"
     :owner "knowledge"
     :description "Local lexical concept index that reduces semantic duplication before code changes.")))

(defparameter *ddd-patterns*
  #("Entity" "ValueObject" "Aggregate" "DomainEvent" "Repository" "Policy"
    "Port" "Adapter" "AntiCorruptionLayer"))
