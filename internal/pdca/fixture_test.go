package pdca

const policyFixture = `{
  "context": "AgentCluster",
  "version": "v1",
  "generated_artifact": "generated/pdca.generated.json",
  "private_cycle_ledger": "data/private/pdca/cycles.jsonl",
  "append_only": true,
  "public_status_redacted": true,
  "raw_cycle_public_allowed": false,
  "steps": [
    {"id":"plan","role":"producer","artifact":"generated/planner.generated.json","command":"mhj planner status"},
    {"id":"do","role":"producer","artifact":"generated/control_plane.generated.json","command":"mhj loop once"},
    {"id":"check","role":"deterministic_verifier","artifact":"generated/verification_evidence.generated.json","command":"mhj verification evidence"},
    {"id":"act","role":"governance_steward","artifact":"generated/learning.generated.json","command":"mhj learning status"}
  ],
  "required_fields": ["cycle_id","at","status","owner","plan_ref","do_ref","check_ref","act_ref"],
  "allowed_statuses": ["open","checking","acting","closed"],
  "evidence_sources": ["LearningLedger","EvidenceGraph","HumanReviewCapacity","AuthorityGate","VerificationEvidenceOps","QualityLedger"],
  "public_summary_fields": ["policy_path","ledger_path","exists","cycle_count","open_count","closed_count","invalid_cycle_count","step_count","ready_step_count","missing_artifact_count","evidence_source_count","by_status","ready","checked_at"],
  "forbidden_public_fields": ["raw_observation","summary","next_action","evidence_ref","evidence_refs","raw_prompt","raw_transcript","token","secret","credential","cookie","account_id","card_number","local_absolute_path","linear_url","private_evidence"],
  "commands": ["mhj pdca status"]
}`
