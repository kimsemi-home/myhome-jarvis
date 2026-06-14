use mhj_command::{plan_for, plan_text};
use mhj_commerce as commerce;
use mhj_finance as finance;

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HarnessCase {
    pub name: &'static str,
    pub command: &'static str,
    pub payload: &'static str,
    pub should_pass: bool,
    pub contains: &'static str,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HarnessResult {
    pub name: &'static str,
    pub passed: bool,
    pub message: String,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HarnessReport {
    pub name: &'static str,
    pub passed: bool,
    pub results: Vec<HarnessResult>,
}

impl HarnessReport {
    pub fn new(name: &'static str, results: Vec<HarnessResult>) -> Self {
        let passed = results.iter().all(|result| result.passed);
        Self {
            name,
            passed,
            results,
        }
    }
}

pub fn home_cases() -> Vec<HarnessCase> {
    vec![
        HarnessCase {
            name: "open_youtube empty payload success",
            command: "open-youtube",
            payload: "{}",
            should_pass: true,
            contains: "https://www.youtube.com",
        },
        HarnessCase {
            name: "open_youtube_search lofi music success",
            command: "open-youtube-search",
            payload: r#"{"query":"lofi music"}"#,
            should_pass: true,
            contains: "search_query=lofi+music",
        },
        HarnessCase {
            name: "open_ott netflix success",
            command: "open-ott",
            payload: r#"{"service":"netflix"}"#,
            should_pass: true,
            contains: "https://www.netflix.com",
        },
        HarnessCase {
            name: "open_netflix shortcut success",
            command: "open-netflix",
            payload: "{}",
            should_pass: true,
            contains: "https://www.netflix.com",
        },
        HarnessCase {
            name: "open_disney_plus shortcut success",
            command: "open-disney-plus",
            payload: "{}",
            should_pass: true,
            contains: "https://www.disneyplus.com",
        },
        HarnessCase {
            name: "open_tving shortcut success",
            command: "open-tving",
            payload: "{}",
            should_pass: true,
            contains: "https://www.tving.com",
        },
        HarnessCase {
            name: "open_wavve shortcut success",
            command: "open-wavve",
            payload: "{}",
            should_pass: true,
            contains: "https://www.wavve.com",
        },
        HarnessCase {
            name: "open_coupang_play shortcut success",
            command: "open-coupang-play",
            payload: "{}",
            should_pass: true,
            contains: "https://www.coupangplay.com",
        },
        HarnessCase {
            name: "open_ott unknown fail",
            command: "open-ott",
            payload: r#"{"service":"unknown"}"#,
            should_pass: false,
            contains: "",
        },
        HarnessCase {
            name: "volume_set 30 success",
            command: "volume-set",
            payload: r#"{"level":30}"#,
            should_pass: true,
            contains: "30",
        },
        HarnessCase {
            name: "volume_set 101 fail",
            command: "volume-set",
            payload: r#"{"level":101}"#,
            should_pass: false,
            contains: "",
        },
        HarnessCase {
            name: "volume_up step 10 success",
            command: "volume-up",
            payload: r#"{"step":10}"#,
            should_pass: true,
            contains: "+ 10",
        },
        HarnessCase {
            name: "volume_down step 10 success",
            command: "volume-down",
            payload: r#"{"step":10}"#,
            should_pass: true,
            contains: "- 10",
        },
        HarnessCase {
            name: "display_sleep success",
            command: "display-sleep",
            payload: "{}",
            should_pass: true,
            contains: "displaysleepnow",
        },
        HarnessCase {
            name: "open_url https success",
            command: "open-url",
            payload: r#"{"url":"https://example.com"}"#,
            should_pass: true,
            contains: "https://example.com",
        },
        HarnessCase {
            name: "open_url javascript fail",
            command: "open-url",
            payload: r#"{"url":"javascript:alert(1)"}"#,
            should_pass: false,
            contains: "",
        },
        HarnessCase {
            name: "movie_mode dry-run success",
            command: "movie-mode",
            payload: "{}",
            should_pass: true,
            contains: "movie_volume",
        },
        HarnessCase {
            name: "sleep_mode dry-run success",
            command: "sleep-mode",
            payload: "{}",
            should_pass: true,
            contains: "display_sleep",
        },
    ]
}

pub fn run_home_harness_report() -> HarnessReport {
    HarnessReport::new("home", run_home_harness())
}

pub fn run_home_harness() -> Vec<HarnessResult> {
    home_cases()
        .into_iter()
        .map(|case| {
            let result = plan_for(case.command, case.payload);
            match (case.should_pass, result) {
                (true, Ok(plan)) => {
                    let text = plan_text(&plan);
                    if case.contains.is_empty() || text.contains(case.contains) {
                        HarnessResult {
                            name: case.name,
                            passed: true,
                            message: "ok".to_string(),
                        }
                    } else {
                        HarnessResult {
                            name: case.name,
                            passed: false,
                            message: format!("missing {}", case.contains),
                        }
                    }
                }
                (true, Err(error)) => HarnessResult {
                    name: case.name,
                    passed: false,
                    message: error.to_string(),
                },
                (false, Ok(_)) => HarnessResult {
                    name: case.name,
                    passed: false,
                    message: "expected failure".to_string(),
                },
                (false, Err(error)) => HarnessResult {
                    name: case.name,
                    passed: true,
                    message: format!("failed safely: {error}"),
                },
            }
        })
        .collect()
}

pub fn run_finance_harness_report() -> HarnessReport {
    HarnessReport::new("finance", run_finance_harness())
}

pub fn run_finance_harness() -> Vec<HarnessResult> {
    let transactions = match finance::fixture_transactions() {
        Ok(transactions) => transactions,
        Err(error) => {
            return vec![check(
                "finance fixture parses",
                false,
                format!("fixture parse failed: {error}"),
            )]
        }
    };
    let summary = match finance::summarize_cashflow(&transactions, "KRW") {
        Ok(summary) => summary,
        Err(error) => {
            return vec![check(
                "finance cashflow summarizes",
                false,
                format!("cashflow failed: {error}"),
            )]
        }
    };
    let owners = match finance::summarize_by_owner(&transactions, "KRW") {
        Ok(owners) => owners,
        Err(error) => {
            return vec![check(
                "finance owners summarize",
                false,
                format!("owner summary failed: {error}"),
            )]
        }
    };
    let subscriptions = match finance::subscription_candidates(&transactions) {
        Ok(subscriptions) => subscriptions,
        Err(error) => {
            return vec![check(
                "subscription candidates summarize",
                false,
                format!("subscription summary failed: {error}"),
            )]
        }
    };
    let cards = match finance::card_usage_candidates(&transactions, "KRW") {
        Ok(cards) => cards,
        Err(error) => {
            return vec![check(
                "card usage candidates summarize",
                false,
                format!("card summary failed: {error}"),
            )]
        }
    };

    let user = finance_owner(&owners, finance::Owner::User);
    let household = finance_owner(&owners, finance::Owner::Household);
    vec![
        check(
            "finance fixture records",
            transactions.len() == 3 && summary.records == 3,
            format!("records={} summary={}", transactions.len(), summary.records),
        ),
        check(
            "finance currency",
            summary.currency == "KRW",
            format!("currency={}", summary.currency),
        ),
        check(
            "finance credit total",
            summary.credit_minor_units == 4_500_000,
            format!("credit={}", summary.credit_minor_units),
        ),
        check(
            "finance debit total",
            summary.debit_minor_units == 153_200,
            format!("debit={}", summary.debit_minor_units),
        ),
        check(
            "finance net total",
            summary.net_minor_units == 4_346_800,
            format!("net={}", summary.net_minor_units),
        ),
        check(
            "subscription review candidates",
            subscriptions.len() == 1
                && summary.subscription_count == 1
                && summary.subscription_minor_units == 65_900
                && subscriptions[0].merchant_name == "Streaming Bundle",
            format!(
                "subscriptions={} total={}",
                summary.subscription_count, summary.subscription_minor_units
            ),
        ),
        check(
            "card-linked debit review candidates",
            cards.len() == 1
                && cards[0].transaction_count == 2
                && cards[0].debit_minor_units == 153_200
                && cards[0].subscription_count == 1,
            cards
                .first()
                .map(|card| {
                    format!(
                        "cards={} transactions={} debit={}",
                        cards.len(),
                        card.transaction_count,
                        card.debit_minor_units
                    )
                })
                .unwrap_or_else(|| "cards=0".to_string()),
        ),
        check(
            "user finance scope",
            user.is_some_and(|owner| owner.records == 1 && owner.net_minor_units == -87_300),
            finance_owner_message(user),
        ),
        check(
            "household finance scope",
            household.is_some_and(|owner| owner.records == 2 && owner.net_minor_units == 4_434_100),
            finance_owner_message(household),
        ),
    ]
}

pub fn run_commerce_harness_report() -> HarnessReport {
    HarnessReport::new("commerce", run_commerce_harness())
}

pub fn run_commerce_harness() -> Vec<HarnessResult> {
    let purchases = match commerce::fixture_purchases() {
        Ok(purchases) => purchases,
        Err(error) => {
            return vec![check(
                "commerce fixture parses",
                false,
                format!("fixture parse failed: {error}"),
            )]
        }
    };
    let summary = match commerce::summarize_commerce(&purchases, "KRW") {
        Ok(summary) => summary,
        Err(error) => {
            return vec![check(
                "commerce summary builds",
                false,
                format!("commerce summary failed: {error}"),
            )]
        }
    };
    let owners = match commerce::summarize_by_owner(&purchases, "KRW") {
        Ok(owners) => owners,
        Err(error) => {
            return vec![check(
                "commerce owners summarize",
                false,
                format!("owner summary failed: {error}"),
            )]
        }
    };
    let merchants = match commerce::summarize_by_merchant(&purchases, "KRW") {
        Ok(merchants) => merchants,
        Err(error) => {
            return vec![check(
                "commerce merchants summarize",
                false,
                format!("merchant summary failed: {error}"),
            )]
        }
    };
    let recurring = match commerce::recurring_candidates(&purchases) {
        Ok(recurring) => recurring,
        Err(error) => {
            return vec![check(
                "recurring purchase candidates summarize",
                false,
                format!("recurring summary failed: {error}"),
            )]
        }
    };

    let user = commerce_owner(&owners, commerce::Owner::User);
    let household = commerce_owner(&owners, commerce::Owner::Household);
    let top_merchant = merchants.first();
    let recurring_candidate = recurring.first();
    vec![
        check(
            "commerce fixture records",
            purchases.len() == 3 && summary.records == 3,
            format!("records={} summary={}", purchases.len(), summary.records),
        ),
        check(
            "commerce currency",
            summary.currency == "KRW",
            format!("currency={}", summary.currency),
        ),
        check(
            "commerce spend total",
            summary.total_spend_minor_units == 26_800,
            format!("spend={}", summary.total_spend_minor_units),
        ),
        check(
            "commerce recurring fixture rows",
            summary.recurring_candidate_count == 2,
            format!("recurring_rows={}", summary.recurring_candidate_count),
        ),
        check(
            "commerce top merchant",
            top_merchant.is_some_and(|merchant| {
                merchant.merchant_name == "Coupang"
                    && merchant.records == 2
                    && merchant.total_spend_minor_units == 23_600
            }),
            top_merchant
                .map(|merchant| {
                    format!(
                        "{} records={} spend={}",
                        merchant.merchant_name, merchant.records, merchant.total_spend_minor_units
                    )
                })
                .unwrap_or_else(|| "missing merchant".to_string()),
        ),
        check(
            "recurring purchase candidate detail",
            recurring_candidate.is_some_and(|candidate| {
                candidate.merchant_name == "Coupang"
                    && candidate.item_name == "Bottled water 2L x 6"
                    && candidate.purchase_count == 2
                    && candidate.latest_total_minor_units == 11_800
            }),
            recurring_candidate
                .map(|candidate| {
                    format!(
                        "{} {} count={} total={}",
                        candidate.merchant_name,
                        candidate.item_name,
                        candidate.purchase_count,
                        candidate.latest_total_minor_units
                    )
                })
                .unwrap_or_else(|| "missing recurring candidate".to_string()),
        ),
        check(
            "user commerce scope",
            user.is_some_and(|owner| owner.records == 1 && owner.total_spend_minor_units == 3_200),
            commerce_owner_message(user),
        ),
        check(
            "household commerce scope",
            household
                .is_some_and(|owner| owner.records == 2 && owner.total_spend_minor_units == 23_600),
            commerce_owner_message(household),
        ),
    ]
}

fn check(name: &'static str, passed: bool, message: impl Into<String>) -> HarnessResult {
    HarnessResult {
        name,
        passed,
        message: if passed {
            "ok".to_string()
        } else {
            message.into()
        },
    }
}

fn finance_owner(
    owners: &[finance::OwnerCashflowSummary],
    owner: finance::Owner,
) -> Option<&finance::OwnerCashflowSummary> {
    owners.iter().find(|summary| summary.owner == owner)
}

fn finance_owner_message(owner: Option<&finance::OwnerCashflowSummary>) -> String {
    owner
        .map(|owner| {
            format!(
                "{} records={} net={}",
                owner.owner.as_str(),
                owner.records,
                owner.net_minor_units
            )
        })
        .unwrap_or_else(|| "missing owner".to_string())
}

fn commerce_owner(
    owners: &[commerce::OwnerSpendSummary],
    owner: commerce::Owner,
) -> Option<&commerce::OwnerSpendSummary> {
    owners.iter().find(|summary| summary.owner == owner)
}

fn commerce_owner_message(owner: Option<&commerce::OwnerSpendSummary>) -> String {
    owner
        .map(|owner| {
            format!(
                "{} records={} spend={}",
                owner.owner.as_str(),
                owner.records,
                owner.total_spend_minor_units
            )
        })
        .unwrap_or_else(|| "missing owner".to_string())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn home_harness_passes() {
        let results = run_home_harness();
        assert!(results.iter().all(|result| result.passed), "{results:#?}");
    }

    #[test]
    fn home_harness_report_passes() {
        let report = run_home_harness_report();
        assert_eq!(report.name, "home");
        assert!(report.passed, "{:#?}", report.results);
        assert_eq!(report.results.len(), home_cases().len());
    }

    #[test]
    fn finance_harness_passes() {
        let report = run_finance_harness_report();
        assert_eq!(report.name, "finance");
        assert!(report.passed, "{:#?}", report.results);
        assert!(report.results.len() >= 9);
    }

    #[test]
    fn commerce_harness_passes() {
        let report = run_commerce_harness_report();
        assert_eq!(report.name, "commerce");
        assert!(report.passed, "{:#?}", report.results);
        assert!(report.results.len() >= 8);
    }
}
