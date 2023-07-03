pub fn reverse(input: &str) -> String {
    let mut res = String::with_capacity(input.len());

    res.extend((*input).chars().rev());

    res
}
