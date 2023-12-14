use std::cmp::{max, min};

fn gcd(a: usize, b: usize) -> usize {
    match ((a, b), (a & 1, b & 1)) {
        ((x, y), _) if x == y => y,
        ((0, x), _) | ((x, 0), _) => x,
        ((x, y), (0, 1)) | ((y, x), (1, 0)) => gcd(x >> 1, y),
        ((x, y), (0, 0)) => gcd(x >> 1, y >> 1) << 1,
        ((x, y), (1, 1)) => {
            let (x, y) = (min(x, y), max(x, y));
            gcd((y - x) >> 1, x)
        }
        _ => unreachable!(),
    }
}

fn lcm(a: usize, b: usize) -> usize {
    a * b / gcd(a, b)
}

type Point = [isize; 3];

#[derive(Debug, Clone, PartialEq)]
struct Moon {
    position: Point,
    velocity: Point,
}

impl Moon {
    fn new(x: isize, y: isize, z: isize) -> Self {
        Moon {
            position: [x, y, z],
            velocity: [0, 0, 0]
        }
    }
}

fn part1(time: isize, moons: &mut [Moon]) -> isize {
    for _step in 1..=time {
        // gravity
        for i in 0..moons.len() {
            for j in 0..moons.len() {
                for dim in 0..3 {
                    moons[j].velocity[dim] += (moons[i].position[dim] - moons[j].position[dim]).signum();
                }
            }
        }

        // velocity
        for i in 0..moons.len() {
            for dim in 0..3 {
                moons[i].position[dim] += moons[i].velocity[dim];
            }
        }
    }

    moons.iter().map(|moon| {
        let mut potential = 0;
        let mut kinetic = 0;

        for dim in 0..3 {
            potential += moon.position[dim].abs();
            kinetic += moon.velocity[dim].abs();
        }

        (potential * kinetic) as isize
    }).sum()
}

fn period(dim: usize, initial_state: &[Moon]) -> usize {
    let mut moons = initial_state.to_vec();

    let mut step = 0;

    loop {
        step += 1;

        // gravity
        for i in 0..moons.len() {
            for j in 0..moons.len() {
                moons[j].velocity[dim] += (moons[i].position[dim] - moons[j].position[dim]).signum();
            }
        }

        // velocity
        for i in 0..moons.len() {
            moons[i].position[dim] += moons[i].velocity[dim];
        }

        if moons == initial_state {
            return step
        }
    }
}

fn part2(moons: &[Moon]) -> usize {
    let x_period = period(0, &moons);
    let y_period = period(1, &moons);
    let z_period = period(2, &moons);

    lcm(x_period, lcm(y_period, z_period))
}

fn main() {
    let mut moons = [
        Moon::new(-6, -5, -8),
        Moon::new(0, -3, -13),
        Moon::new(-15, 10, -11),
        Moon::new(-3, -8, 3)
    ];

    println!("Part 1: {}", part1(1000, &mut moons));
    println!("Part 2: {:?}", part2(&moons));
}
