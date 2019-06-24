import subprocess
import os

inputs  = [ f'inputs/{each}' for each in os.listdir('inputs') ]
for each in inputs:
    of = each.replace('input', 'output')
    subprocess.run(['./sudoku', each, of])