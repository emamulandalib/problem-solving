use std::io;

pub fn bit_plus_plus() {
    let mut nums_of_statement_inp = String::new();
    io::stdin().read_line(&mut nums_of_statement_inp).expect("TODO: panic message");
    let nums_of_statement = nums_of_statement_inp.trim().parse::<i32>().expect("Failed to parse");

    let mut result = 0;

    for _ in 0..nums_of_statement {
        let mut statement_inp = String::new();
        io::stdin().read_line(&mut statement_inp).expect("TODO: failed to read statement");
        let statement = statement_inp.trim();

        match statement {
            "X++" | "++X" => result += 1,
            "X--" | "--X" => result -= 1,
            _ => {}
        }
    }

    println!("{}", result)
}