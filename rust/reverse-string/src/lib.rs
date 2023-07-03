pub fn reverse(input: &str) -> String {
    let mut res = String::with_capacity(input.len());

    for chr in (*input).chars().rev() {
        res.push(chr);
    }

    res.to_string()
}
