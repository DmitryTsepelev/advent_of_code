use std::fs::File;
use std::io::BufReader;
use std::io::BufRead;
use std::collections::HashSet;
use std::collections::HashMap;

#[derive(Debug)]
struct Leg {
    direction: char,
    length: i32
}

impl Leg {
    fn new(source: &str) -> Leg {
        let direction = &source.chars().next().unwrap();
        let length: i32 = source[1..source.len()].parse().expect("Failed parse number!");

        Leg { direction: *direction, length: length }
    }
}

#[derive(Debug, Hash)]
struct Point {
    x: i32,
    y: i32
}

impl PartialEq for Point {
    fn eq(&self, other: &Self) -> bool {
        self.x == other.x && self.y == other.y
    }
}

impl Eq for Point {}

fn part1(first_wire: Vec<Leg>, second_wire: Vec<Leg>) {
    let mut points = HashSet::new();

    let mut current_x = 0;
    let mut current_y = 0;
    for leg in first_wire {
        let mut dx = 0;
        let mut dy = 0;

        match leg.direction {
            'U' => dy = -1,
            'D' => dy = 1,
            'L' => dx = -1,
            'R' => dx = 1,
            _ => ()
        }

        for _i in 0..leg.length {
            current_x += dx;
            current_y += dy;

            let point = Point { x: current_x, y: current_y };
            points.insert(point);
        }
    }

    let mut min_distance = std::i32::MAX;

    let mut current_x = 0;
    let mut current_y = 0;
    for leg in second_wire {
        let mut dx = 0;
        let mut dy = 0;

        match leg.direction {
            'U' => dy = -1,
            'D' => dy = 1,
            'L' => dx = -1,
            'R' => dx = 1,
            _ => ()
        }

        for _i in 0..leg.length {
            current_x += dx;
            current_y += dy;

            let point = Point { x: current_x, y: current_y };

            if points.contains(&point) {
                let distance = i32::abs(point.x) + i32::abs(point.y);
                if distance < min_distance {
                    min_distance = distance;
                }
            }
        }
    }

    println!("Part 1 (min_distance): {}", min_distance)
}

fn part2(first_wire: Vec<Leg>, second_wire: Vec<Leg>) {
    let mut points = HashMap::new();

    let mut current_x = 0;
    let mut current_y = 0;
    let mut path = 0;
    for leg in first_wire {
        let mut dx = 0;
        let mut dy = 0;

        match leg.direction {
            'U' => dy = -1,
            'D' => dy = 1,
            'L' => dx = -1,
            'R' => dx = 1,
            _ => ()
        }

        for _i in 0..leg.length {
            current_x += dx;
            current_y += dy;
            path += 1;

            let point = Point { x: current_x, y: current_y };
            points.insert(point, path);
        }
    }

    let mut min_path = std::i32::MAX;

    let mut current_x = 0;
    let mut current_y = 0;
    let mut path = 0;
    for leg in second_wire {
        let mut dx = 0;
        let mut dy = 0;

        match leg.direction {
            'U' => dy = -1,
            'D' => dy = 1,
            'L' => dx = -1,
            'R' => dx = 1,
            _ => ()
        }

        for _i in 0..leg.length {
            current_x += dx;
            current_y += dy;
            path += 1;

            let point = Point { x: current_x, y: current_y };

            match points.get(&point) {
                Some(first_wire_path) => {
                    let total_path = first_wire_path + path;
                    if total_path < min_path {
                        min_path = total_path;
                    }
                },
                None => ()
            }
        }
    }

    println!("Part 2 (min_path): {}", min_path)
}

fn main() {
    let filename = "input.txt".to_string();
    let file = File::open(filename).expect("Cannot read file!");
    let reader = BufReader::new(file);

    let lines: Vec<String> = reader.lines().map(|line| {
        let line = line.expect("Failed reading line!").parse().expect("Failed parse line!");
        line
    }).collect();

    let first_wire: Vec<Leg> = lines[0].split(",").map(|line| Leg::new(line)).collect();
    let second_wire: Vec<Leg> = lines[1].split(",").map(|line| Leg::new(line)).collect();

    part1(first_wire, second_wire);
    part2(first_wire, second_wire);
}
