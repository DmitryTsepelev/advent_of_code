use std::fs::File;
use std::io::BufReader;
use std::io::BufRead;
use std::collections::HashMap;

fn orbit_to_com(orbits: &HashMap<String, String>, start: &str) -> i32 {
    let mut count = 0;

    let mut current = start;

    loop {
        match orbits.get(current) {
            Some(object) => {
                current = object;
                count += 1;
            },
            None => break,
        };
    }

    count
}

fn part1(orbits: &HashMap<String, String>) {
    let count: i32 = orbits.keys().map(|child| orbit_to_com(&orbits, child)).sum();

    println!("Part 1 = {}", count);
}

fn build_path(orbits: &HashMap<String, String>, start: &str) -> Vec<String> {
    let mut current = &orbits[start];
    let mut path = vec![current.to_string()];

    loop {
        match orbits.get(current) {
            Some(object) => {
                current = object;
                path.push(object.to_string());
            },
            None => break,
        };
    }

    path.reverse();
    path
}

fn part2(orbits: &HashMap<String, String>) {
    let to_com = orbit_to_com(orbits, &orbits["YOU"]) + orbit_to_com(orbits, &orbits["SAN"]);

    let mut common_length = 0;
    for (x, y) in build_path(orbits, "YOU").iter().zip(build_path(orbits, "SAN")) {
        if x.to_string() == y.to_string() {
            common_length += 1;
        } else {
            break;
        }
    }

    let result = to_com - (common_length - 1) * 2;

    println!("Part 2 = {}", result);
}

fn main() {
    let filename = "input.txt".to_string();
    let file = File::open(filename).expect("Cannot read file!");
    let reader = BufReader::new(file);

    let mut orbits = HashMap::new();

    for line in reader.lines() {
        let line = line.unwrap();
        let objects: Vec<&str> = line.split(")").collect();

        orbits.insert(objects[1].to_string(), objects[0].to_string());
    }

    part1(&orbits);
    part2(&orbits);
}
