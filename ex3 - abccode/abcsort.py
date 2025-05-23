# Импортируем только нужную функцию из collections
from collections import Counter

# Функция для чтения чисел из файла
def read_numbers(filename):
    with open(filename) as f:
        return [int(num) for num in f.read().split()]

# Читаем числа из обоих файлов
nums_a = read_numbers('A.txt')
nums_b = read_numbers('B.txt')

# Считаем количество каждого числа в обоих файлах
counter_a = Counter(nums_a)
counter_b = Counter(nums_b)

# Находим общие числа для обоих файлов
common_nums = set(nums_a).intersection(set(nums_b))

# Формируем результат: для каждого общего числа берем максимальное количество из обоих файлов
output = []
for num in common_nums:
    max_count = max(counter_a[num], counter_b[num])
    output += [num] * max_count

# Сортируем результат и записываем в файл
output.sort()
with open('C.txt', 'w') as f:
    f.write('\n'.join(map(str, output)))