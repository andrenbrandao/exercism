use std::{ops::Add, time::Duration};

use time::PrimitiveDateTime as DateTime;

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    return start.add(Duration::new(1000000000, 0));
}
