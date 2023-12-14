use std::fs::File;
use std::io::BufReader;
use std::io::BufRead;

fn fuel_simple(mass: f32) -> f32 {
    (mass / 3.0).floor() - 2.0
}

fn fuel_advanced(mass: f32) -> f32 {
    let base = fuel_simple(mass);

    if base <= 0.0 {
        0.0
    } else {
        base + fuel_advanced(base)
    }
}

fn solve(advanced: bool) {
    let filename = "input.txt".to_string();
    let file = File::open(filename).expect("Cannot read file!");
    let reader = BufReader::new(file);

    let fuel: f32 = reader.lines().map(|line| {
        let line = line.expect("Failed reading line!");
        let mass: f32 = line.parse().expect("Failed parse number!");

        if advanced {
            fuel_advanced(mass)
        } else {
            fuel_simple(mass)
        }
    }).sum();

    let mut advanced_text = "".to_string();
    if advanced {
        advanced_text = " (advanced)".to_string();
    }

    println!("Fuel{} = {}", advanced_text, fuel);
}

fn main() {
    solve(false);
    solve(true);
}
