use time::{Duration, PrimitiveDateTime};

// Returns a DateTime one billion seconds after start.
pub fn after(start: PrimitiveDateTime) -> PrimitiveDateTime {
    const GIGASECOND: Duration = Duration::new(1_000_000_000, 0);
    start + GIGASECOND
}
