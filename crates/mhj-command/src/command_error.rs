use std::error::Error;
use std::fmt;

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum CommandError {
    UnknownCommand(String),
    InvalidPayload(String),
}

impl fmt::Display for CommandError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            CommandError::UnknownCommand(name) => write!(formatter, "unknown command {name}"),
            CommandError::InvalidPayload(message) => {
                write!(formatter, "invalid payload: {message}")
            }
        }
    }
}

impl Error for CommandError {}
