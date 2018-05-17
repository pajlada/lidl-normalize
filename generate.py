#!/usr/bin/env python3

from collections import OrderedDict

import traceback

unicodeFileIn = 'files/confusables.txt'
customFileIn = 'files/custom-confusables.txt'

confusables = OrderedDict()

def insertConfusablesFromFile(path):
        with open(path, 'r') as fh:
            for line in fh.readlines():
                if len(line) == 0:
                    continue

                if line[0] == '#':
                    continue

                parts = line.split(';')
                if len(parts) < 2:
                    continue

                inPart = parts[0].strip().lower()
                outPart = parts[1].strip().lower()
                confusables[inPart] = outPart

def insertConfusablesFromFile2(path):
        with open(path, 'r') as fh:
            for line in fh.readlines():
                if len(line) == 0:
                    continue

                if line[0] == '#':
                    continue

                parts = line.split(':')
                if len(parts) != 2:
                    continue

                inPart = parts[0].strip().lower()[2:]
                outPart = hex(ord(parts[1].strip()[1]))[2:]
                confusables[inPart] = outPart

def insertConfusablesFromFile3(path):
        with open(path, 'r') as fh:
            for line in fh.readlines():
                if len(line) == 0:
                    continue

                if line[0] == '#':
                    continue

                line = line.replace('"', '')
                line = line.replace(',', '')
                line = line.replace('\'', '')

                parts = line.split(':')
                if len(parts) != 2:
                    continue

                try:
                    inPart = hex(ord(parts[0].strip()))[2:]
                except:
                    print('forsenT {}'.format(parts[0]))
                    print(traceback.format_exc())
                    continue
                outPart = ' '.join(hex(ord(c))[2:] for c in parts[1].strip())
                confusables[inPart] = outPart

insertConfusablesFromFile(unicodeFileIn)
insertConfusablesFromFile(customFileIn)
insertConfusablesFromFile2('files/custom-confusables2.txt')
insertConfusablesFromFile3('files/custom-confusables3.txt')

print('Generating table.go with {} confusables'.format(len(confusables)))

filestring = """package normalize

var confusableTable = map[rune][]rune{}

func init() {
"""

for confusable in confusables:
    simplified = confusables[confusable]
    filestring += 'confusableTable[0x' + confusable + '] = []rune{'
    filestring += ','.join('0x' + s for s in simplified.split(' '))
    filestring += '}\n'

filestring += """
}"""

# print(filestring)

with open('table.go', 'w') as fh:
    fh.write(filestring)
