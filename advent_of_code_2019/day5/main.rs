use std::fs;
use std::str::FromStr;

const ADD_OPCODE: i32 = 1;
const MULT_OPCODE: i32 = 2;
const SAVE_OPCODE: i32 = 3;
const OUTPUT_OPCODE: i32 = 4;
const JUMP_IF_TRUE: i32 = 5;
const JUMP_IF_FALSE: i32 = 6;
const LESS_THAN: i32 = 7;
const EQUALS: i32 = 8;
const EXIT_OPCODE: i32 = 99;

const POSITIONAL_MODE: i32 = 0;
const IMMEDIATE_MODE: i32 = 1;

struct VM<'a> {
    memory: &'a mut [i32],
    pointer: usize,
    input: i32,
    output: i32
}

impl VM<'_> {
    fn get_param(&self, shift: usize) -> i32 {
        let pointer = self.pointer + shift;
        let mode = self.mode_for(shift - 1);

        match mode {
            POSITIONAL_MODE => {
                let position = self.memory[pointer] as usize;
                self.memory[position]
            },
            IMMEDIATE_MODE => self.memory[pointer],
            _ => -1
        }
    }

    fn get_address(&self, shift: usize) -> usize {
       self.memory[self.pointer + shift] as usize
    }

    fn add(&mut self) {
        let first_param = self.get_param(1);
        let second_param = self.get_param(2);
        let index = self.get_address(3);
        self.memory[index] = first_param + second_param;

        self.pointer += 4;
    }

    fn mult(&mut self) {
        let first_param = self.get_param(1);
        let second_param = self.get_param(2);
        let index = self.get_address(3);

        self.memory[index] = first_param * second_param;

        self.pointer += 4;
    }

    fn save(&mut self) {
        let index = self.get_address(1);
        self.memory[index] = self.input;
        self.pointer += 2;
    }

    fn output(&mut self) {
        let index = self.get_address(1);
        self.output = self.memory[index];
        self.pointer += 2;
    }

    fn jump_if_true(&mut self) {
        let first_param = self.get_param(1);

        if first_param != 0 {
            let second_param = self.get_param(2);
            self.pointer = second_param as usize;
        } else {
            self.pointer += 3;
        }
    }

    fn jump_if_false(&mut self) {
        let first_param = self.get_param(1);

        if first_param == 0 {
            let second_param = self.get_param(2);
            self.pointer = second_param as usize;
        } else {
            self.pointer += 3;
        }
    }

    fn less_than(&mut self) {
        let first_param = self.get_param(1);
        let second_param = self.get_param(2);
        let index = self.get_address(3);

        self.memory[index] = if first_param < second_param { 1 } else { 0 };

        self.pointer += 4;
    }

    fn equals(&mut self) {
        let first_param = self.get_param(1);
        let second_param = self.get_param(2);
        let index = self.get_address(3);

        self.memory[index] = if first_param == second_param { 1 } else { 0 };

        self.pointer += 4;
    }

    fn current_opcode(&self) -> i32 {
        self.memory[self.pointer] % 100
    }

    fn mode_for(&self, param_index: usize) -> i32 {
        let mut mode = self.memory[self.pointer] / 100;

        for _i in 0..param_index {
            mode = mode / 10;
        }

        mode % 10
    }

    fn execute(&mut self) {
        while self.current_opcode() != EXIT_OPCODE {
            match self.current_opcode() {
                ADD_OPCODE => self.add(),
                MULT_OPCODE => self.mult(),
                SAVE_OPCODE => self.save(),
                OUTPUT_OPCODE => self.output(),
                JUMP_IF_TRUE => self.jump_if_true(),
                JUMP_IF_FALSE => self.jump_if_false(),
                LESS_THAN => self.less_than(),
                EQUALS => self.equals(),
                _ => ()
            }
        }

        println!("output: {}", self.output);
    }
}

fn main() {
    let contents = fs::read_to_string("input.txt")
        .expect("Something went wrong reading the file")
        .replace("\n", "");

    println!("Part 1:");

    let mut memory: Vec<i32> = contents.split(",").map(|x| i32::from_str(x).unwrap()).collect();

    let mut vm = VM {
        input: 1,
        output: 0,
        memory: memory.as_mut_slice(),
        pointer: 0
    };

    vm.execute();

    println!("Part 2:");

    let mut memory: Vec<i32> = contents.split(",").map(|x| i32::from_str(x).unwrap()).collect();

    let mut vm = VM {
        input: 5,
        output: 0,
        memory: memory.as_mut_slice(),
        pointer: 0
    };

    vm.execute();
}
