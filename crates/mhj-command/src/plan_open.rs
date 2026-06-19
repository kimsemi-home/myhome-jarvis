use crate::json::json_string;
use crate::ott::ott_url;
use crate::strings::open_url_plan;
use crate::{CommandError, Plan};

pub(crate) fn open_ott_plan(payload: &str) -> Result<Plan, CommandError> {
    let service = json_string(payload, "service")?.to_ascii_lowercase();
    let url = ott_url(&service).ok_or_else(|| {
        CommandError::InvalidPayload(format!("unsupported ott service {service}"))
    })?;
    Ok(open_url_plan("open_ott", url))
}

pub(crate) fn normalize_command(command: &str) -> String {
    command.trim().to_ascii_lowercase().replace('-', "_")
}
