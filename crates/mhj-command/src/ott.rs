use crate::strings::open_url_plan;
use crate::{CommandError, Plan};

pub(crate) fn ott_shortcut_plan(name: &str, service: &str) -> Result<Plan, CommandError> {
    let target = ott_url(service).ok_or_else(|| {
        CommandError::InvalidPayload(format!("missing OTT shortcut target for {service}"))
    })?;
    Ok(open_url_plan(name, target))
}

pub(crate) fn ott_url(service: &str) -> Option<&'static str> {
    match service {
        "coupangplay" => Some("https://www.coupangplay.com"),
        "disney" => Some("https://www.disneyplus.com"),
        "netflix" => Some("https://www.netflix.com"),
        "tving" => Some("https://www.tving.com"),
        "wavve" => Some("https://www.wavve.com"),
        "youtube" => Some("https://www.youtube.com"),
        _ => None,
    }
}
