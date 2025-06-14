# Отрезки

Дано n отрезков. Каждый из отрезков задается координатами своих концов. Требуется реализовать функцию, которая, получая на вход множество отрезков, определяет сколько нужно взять точек так, чтобы каждому отрезку принадлежала хотя бы одна точка.


## Алгоритм решения

Используется жадный подход:

2. Сортируем массив отрезков по возрастанию правой границы $end$.
3. Инициализируем счётчик точек `count = 1` и ставим первую точку в правой границе первого отрезка: `currentPoint = segments[0].End`.
4. Для каждого последующего отрезка:

   * Если `currentPoint < segments[i].Start`, то текущая точка не покрывает новый отрезок: ставим новую точку `currentPoint = segments[i].End` и увеличиваем `count += 1`.
   * Иначе отрезок уже покрыт — пропускаем.
5. По окончании прохода по всем отрезкам `count` — искомое минимальное число точек.

Время работы: $O(n \log n)$ из-за сортировки.

## Формат ввода

```
<n>
start_1 end_1
start_2 end_2
...
start_n end_n
```

* `n` — количество отрезков.
* Далее `n` строк с двумя целыми числами — координатами концов каждого отрезка.

## Пример

**Вход:**

```
3
1 3
2 5
3 6
```

**Выход:**

```
1
```