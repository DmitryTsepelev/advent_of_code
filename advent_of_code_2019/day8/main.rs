use std::collections::HashMap;
use std::fs;

const WIDTH: usize = 25;
const HEIGHT: usize = 6;

fn part1() {
    let image = fs::read_to_string("input.txt")
        .expect("Something went wrong reading the file");

    let mut offset = 0;
    let mut min_zeros = WIDTH * HEIGHT + 1;
    let mut result = WIDTH * HEIGHT + 1;
    loop {
        if offset == image.len() - 1 {
            break;
        }
        let layer = &image[offset..offset + WIDTH * HEIGHT];

        let counts = layer.chars().fold(HashMap::new(), |mut counts, c| {
            let counter = counts.entry(c).or_insert(0);
            *counter += 1;

            counts
        });

        let zeros = match counts.get(&'0') {
            Some(value) => *value,
            None => 0
        };

        if zeros < min_zeros {
            let ones = match counts.get(&'1') {
                Some(value) => value,
                None => &0
            };

            let twos = match counts.get(&'2') {
                Some(value) => value,
                None => &0
            };

            min_zeros = zeros;
            result = ones * twos;
        }

        offset += WIDTH * HEIGHT;
    }

    println!("Part 1: {}", result);
}

fn part2() {
    let image = fs::read_to_string("input.txt")
        .expect("Something went wrong reading the file");

    let mut decoded = String::new();
    let mut offset = 0;
    loop {
        if offset == image.len() - 1 {
            break;
        }
        let layer = &image[offset..offset + WIDTH * HEIGHT];

        if decoded.len() == 0 {
            decoded = layer.to_string();
        } else {
            let mut new_decoded = String::new();

            for (i, layer_pixel) in layer.chars().enumerate() {
                let decoded_pixel = decoded.chars().nth(i).unwrap();

                if decoded_pixel == '2' {
                    new_decoded.push(layer_pixel);
                } else {
                    new_decoded.push(decoded_pixel);
                }
            }
            decoded = new_decoded;
        }

        offset += WIDTH * HEIGHT;
    }

    let mut offset = 0;
    loop {
        if offset == decoded.len() {
            break;
        }
        let row = &decoded[offset..offset + WIDTH];

        for c in row.chars() {
            match c {
                '0' => print!(" "),
                '1' => print!("|"),
                _ => ()
            }
        }

        println!("");
        offset += WIDTH;
    }
}

fn main() {
    part1();
    part2();
}
