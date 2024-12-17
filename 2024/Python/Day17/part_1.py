import os

os.system("cls")


class ThreeBitMachine:
    def __init__(
        self, register_a: int, register_b: int, register_c: int, program: list[int]
    ):
        self.register_a = register_a
        self.register_b = register_b
        self.register_c = register_c
        self.program = program

        self.instruction_pointer = 0
        self.output_values = []

    def _adv(self, operand: int):
        numerator = self.register_a
        denominator = 2 ** self.get_combo_operand_value(operand=operand)
        self.register_a = numerator // denominator
        self.increment_pointer()

    def _bdv(self, operand: int):
        numerator = self.register_a
        denominator = 2 ** self.get_combo_operand_value(operand=operand)
        self.register_b = numerator // denominator
        self.increment_pointer()

    def _cdv(self, operand: int):
        numerator = self.register_a
        denominator = 2 ** self.get_combo_operand_value(operand=operand)
        self.register_c = numerator // denominator
        self.increment_pointer()

    def _bxl(self, operand: int):
        self.register_b = self.register_b ^ operand
        self.increment_pointer()

    def _bst(self, operand: int):
        self.register_b = self.get_combo_operand_value(operand=operand) % 8
        self.increment_pointer()

    def _jnz(self, operand: int):
        if self.register_a == 0:
            self.increment_pointer()
            return

        self.instruction_pointer = operand

    def _bxc(self, operand: int):
        self.register_b = self.register_b ^ self.register_c
        self.increment_pointer()

    def _out(self, operand: int):
        operand_value = self.get_combo_operand_value(operand)
        self.output_values.append(operand_value % 8)
        self.increment_pointer()

    def increment_pointer(self):
        self.instruction_pointer += 2

    def get_combo_operand_value(self, operand: int):
        if 0 <= operand <= 3:
            return operand

        if operand == 4:
            return self.register_a

        if operand == 5:
            return self.register_b

        if operand == 6:
            return self.register_c

        raise Exception("not a valid program")

    def run_instruction(self, opcode: int, operand: int):
        if opcode == 0:
            return self._adv(operand)

        if opcode == 6:
            return self._bdv(operand)

        if opcode == 7:
            return self._cdv(operand)

        if opcode == 1:
            return self._bxl(operand)

        if opcode == 2:
            return self._bst(operand)

        if opcode == 3:
            return self._jnz(operand)

        if opcode == 4:
            return self._bxc(operand)

        if opcode == 5:
            return self._out(operand)

        raise Exception("not a valid opcode")

    def get_output(self):
        print(",".join([str(value) for value in self.output_values]))

    def run_program(self):
        while True:
            if self.instruction_pointer >= len(self.program):
                break

            if self.instruction_pointer + 1 >= len(self.program):
                break

            opcode = self.program[self.instruction_pointer]
            operand = self.program[self.instruction_pointer + 1]
            self.run_instruction(opcode=opcode, operand=operand)

        self.get_output()


def read_input_file(file_path: str) -> list[str]:
    with open(file=file_path, mode="r") as input_file:
        lines = input_file.readlines()
        return [line.strip() for line in lines]


def solution(lines: list[str]):
    register_a = int(lines[0][12:])
    register_b = int(lines[1][12:])
    register_c = int(lines[2][12:])
    program = list(map(int, lines[4][9:].split(",")))

    machine = ThreeBitMachine(
        register_a=register_a,
        register_b=register_b,
        register_c=register_c,
        program=program,
    )

    machine.run_program()


lines = read_input_file(file_path="input.txt")
solution(lines)
