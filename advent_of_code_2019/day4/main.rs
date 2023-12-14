fn part1() {
    let mut password_count = 0;

    for i in 108457..=562041 {
        let password = i.to_string();

        let mut ascending = true;
        let mut same_digits = false;
        let mut prev_char = '\0';

        for current_char in password.chars() {
            if prev_char > current_char {
                ascending = false;
                break;
            }

            if prev_char == current_char {
                same_digits = true;
            }

            prev_char = current_char;
        }

        if ascending && same_digits {
            password_count += 1;
        }
    }

    println!("Part 1: {}", password_count);
}

fn part2() {
    let mut password_count = 0;

    for i in 108457..=562041 {
        let password = i.to_string();

        let mut ascending = true;
        let mut same_two_digits = false;
        let mut same_digits_count = 0;
        let mut prev_char = '\0';

        for current_char in password.chars() {
            if prev_char > current_char {
                ascending = false;
                break;
            }

            if prev_char == current_char {
                same_digits_count += 1;
            } else {
                if same_digits_count == 2 {
                    same_two_digits = true;
                }
                same_digits_count = 1;
            }

            prev_char = current_char;
        }

        if same_digits_count == 2 {
            same_two_digits = true;
        }

        if ascending && same_two_digits {
            password_count += 1;
        }
    }

    println!("Part 2: {}", password_count);
}

fn main() {
    part1();
    part2();
}
