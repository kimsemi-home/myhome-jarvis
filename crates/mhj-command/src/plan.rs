use crate::json::{json_i32, json_string};
use crate::modes::{display_sleep_plan, mac_sleep_plan, movie_mode_plan, sleep_mode_plan};
use crate::ott::ott_shortcut_plan;
use crate::plan_open::{normalize_command, open_ott_plan};
use crate::strings::open_url_plan;
use crate::url::{encode_query, validate_http_url};
use crate::volume::{self, volume_set, volume_step_plan};
use crate::{CommandError, Plan};

pub fn plan_for(command: &str, payload: &str) -> Result<Plan, CommandError> {
    match normalize_command(command).as_str() {
        "open_coupang_play" => ott_shortcut_plan("open_coupang_play", "coupangplay"),
        "open_disney_plus" => ott_shortcut_plan("open_disney_plus", "disney"),
        "open_netflix" => ott_shortcut_plan("open_netflix", "netflix"),
        "open_youtube" => Ok(open_url_plan("open_youtube", "https://www.youtube.com")),
        "open_youtube_search" => youtube_search_plan(payload),
        "open_tving" => ott_shortcut_plan("open_tving", "tving"),
        "open_wavve" => ott_shortcut_plan("open_wavve", "wavve"),
        "open_ott" => open_ott_plan(payload),
        "open_url" => Ok(open_url_plan(
            "open_url",
            &validate_http_url(&json_string(payload, "url")?)?,
        )),
        "volume_set" => volume_set(json_i32(payload, "level")?),
        "volume_up" => volume_step_plan("volume_up", payload, '+'),
        "volume_down" => volume_step_plan("volume_down", payload, '-'),
        "volume_mute" => volume::mute_plan(),
        "display_sleep" => Ok(display_sleep_plan()),
        "mac_sleep" => Ok(mac_sleep_plan()),
        "movie_mode" => Ok(movie_mode_plan()),
        "sleep_mode" => Ok(sleep_mode_plan()),
        other => Err(CommandError::UnknownCommand(other.to_string())),
    }
}

pub fn plan_text(plan: &Plan) -> String {
    let mut text = plan.name.clone();
    for invocation in &plan.invocations {
        text.push(' ');
        text.push_str(&invocation.label);
        if let Some(url) = &invocation.url {
            text.push(' ');
            text.push_str(url);
        }
        for arg in &invocation.argv {
            text.push(' ');
            text.push_str(arg);
        }
    }
    text
}

fn youtube_search_plan(payload: &str) -> Result<Plan, CommandError> {
    let query = json_string(payload, "query")?;
    if query.trim().is_empty() {
        return Err(CommandError::InvalidPayload(
            "query is required".to_string(),
        ));
    }
    Ok(open_url_plan(
        "open_youtube_search",
        &format!(
            "https://www.youtube.com/results?search_query={}",
            encode_query(query.trim())
        ),
    ))
}
