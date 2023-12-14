use std::fs;
use std::str::FromStr;

const ADD_OPCODE: i32 = 1;
const MULT_OPCODE: i32 = 2;
const EXIT_OPCODE: i32 = 99;

fn execute(memory: &mut Vec<i32>, pointer: usize) {
    let opcode = memory[pointer];

    if opcode == EXIT_OPCODE {
        return
    }

    let index = memory[pointer + 3] as usize;
    let first_num = memory[memory[pointer + 1] as usize];
    let second_num = memory[memory[pointer + 2] as usize];

    if opcode == ADD_OPCODE {
        memory[index] = first_num + second_num;
    } else if opcode == MULT_OPCODE {
        memory[index] = first_num * second_num;
    }

    execute(memory, pointer + 4);
}

fn execute_with(memory: &mut Vec<i32>, verb: i32, noun: i32) -> i32 {
    memory[1] = noun;
    memory[2] = verb;
    execute(memory, 0);
    memory[0]
}

fn main() {
    let contents = fs::read_to_string("input.txt")
        .expect("Something went wrong reading the file")
        .replace("\n", "");

    let mut memory: Vec<i32> = contents.split(",").map(|x| i32::from_str(x).unwrap()).collect();
    let result = execute_with(&mut memory, 2, 12);
    println!("Part 1: {}", result);

    for verb in 0..=99 {
        for noun in 0..=99 {
            let mut memory: Vec<i32> = contents.split(",").map(|x| i32::from_str(x).unwrap()).collect();
            let result = execute_with(&mut memory, verb, noun);

            if result == 19690720 {
                println!("Part 2: {}", 100 * noun + verb);
                return
            }
        }
    }
}
