Реализована программа с использованием каналов, где входные параметры - количество рабочих.
Имеются рабочие, которые выполняют работу параллельно (время работы - случайное число от 1 до 3 секунд),
затем они завершают работу и сдают результаты.

Я не создавал сущности "Прораб". Всё взаимодействие осуществляется через функцию Work,
которая в качестве аргументов принимает количество работников, а возвращает отчёт о проделанной работе.
В роли работы здесь выступает возведение числа в квадрат.

Также добавлен benchmark-test к функции Work.