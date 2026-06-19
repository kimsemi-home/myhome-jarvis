#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Invocation {
    pub label: String,
    pub argv: Vec<String>,
    pub url: Option<String>,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Plan {
    pub name: String,
    pub dry_run: bool,
    pub invocations: Vec<Invocation>,
}
