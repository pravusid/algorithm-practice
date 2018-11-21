# 정렬

## 정렬의 개념

정렬은 여러 원소로 구성된 데이터를 크기 순서에 따라 재배열하는 것

비교 기반 정렬은 데이터의 키 값을 비교하면서 정렬하는 방식으로, 시간 복잡도는 키의 비교횟수에 비례한다.
> 선택 정렬, 버블 정렬, 삽입 정렬, 셸 정렬, 합병 정렬, 퀵 정렬, 힙 정렬 ...

분포 기반 정렬은 데이터의 분포 정보를 이용하여 정렬하는 방식으로, 선형시간 복잡도를 가진다.
> 계수 정렬, 버킷 정렬, 기수 정렬 ...

안정적 정렬은 데이터의 키 값이 동일한 경우 정렬 후 기존 순서가 변하지 않는 정렬이다.
제자리 정렬은 입력받은 데이터가 저장된 공간 이외의 별도 공간을 상수개만 사용하는 정렬이다.

## 선택 정렬

선택정렬은 원소중 가장 작은 키값을 갖는 원소를 선택하여 차례대로 나열하는 정렬방식이다

`{ 30, 50, 7, 40, 88, 15, 44 }` 의 배열을 정렬해보면

1. 첫 번째 원소 30을 기준으로 오른쪽으로 최소값을 찾기 시작한다.
2. 50은 30보다 더 크므로 다시 오른쪽 값을 비교한다.
3. 7은 30보다 더 작으므로 7을 최소값 기준으로 설정한다
4. 40은 7보다 크므로 다시 오른쪽 값을 비교한다.
5. ... 마지막 원소인 44까지 진행한다
6. 44는 7보다 크므로 최소값은 7이된다.
7. 첫 번째 원소와 찾은 최소값인 7의 자리를 바꾼다.
8. 두 번째 원소를 대상으로 1 ~ 7 과정을 반복한다.

### 선택 정렬 알고리즘

```python
def SelectionSort(dlist):
    size = len(dlist)
    for element in range(size): # range는 반폐구간
        min_position = element
        for target in range(element + 1, size):
            if (dlist[target] < dlist[min_position]):
                min_position = target
        (dlist[element], dlist[min_position]) = (dlist[min_position], dlist[element])
```

### 선택 정렬 알고리즘 성능

선택 정렬 알고리즘은 이중 반복문으로 구성되며 둘 다 입력 배열 크기 `n`에 비례하므로 전체 수행시간은 `O(n^2)` 이다.
입력 배열 크기 이외의 인자가 없으므로 최악, 최선, 평균 수형시간 모두 `O(n^2)` 이다.

### 선택 정렬 특징

1. 언제나 동일 수행시간 `O(n^2)`을 보임
2. 안정적이지 않은 정렬임
3. 제자리 정렬임: 데이터가 움직이는 경우는 두 원소 교환밖에 없음 (저장공간 이외의 공간을 상수개만 사용)

## 버블 정렬

버블 정렬은 인접한 두 원소를 차례대로 비교하면서 자리바꿈을 통해 정렬하는 방식이다.

`{ 30, 50, 7, 40, 88, 15, 44 }` 의 배열을 정렬해보면

1. 첫 번째 30과 두 번째 50을 비교한다.
2. 30이 50보다 작다. 두 번째 50과 세 번째 7을 비교한다.
3. 50이 7보다 크다. 두 수의 자리를 바꾼다.
4. 세 번째 50을 다시 네 번째 40과 비교한다.
5. 50이 40보다 크다. 두 수의 자리를 바꾼다.
6. ... 반복하여 n-1번째 원소와 n번째 원소를 비교할때 까지 진행한다.
7. 다시 1 ~ 6의 과정을, 정렬된 원소 이전까지 반복한다.

### 버블 정렬 알고리즘

```python
def BubbleSort(dlist):
    size = len(dlist) - 1
    for element in range(size):
        for head in range(size - element):
            if (dlist[head] > dlist[head + 1]):
                (dlist[head], dlist[head + 1]) = (dlist[head + 1], dlist[head])
```

### 버블 정렬 알고리즘 성능

버블 정렬 알고리즘은 이중 반복문으로 구성되며
바깥 반복은 조건에 따라 최소 상수, 최대 `n`에 비례하고, 안쪽 반복은 `n`에 비례한다.

수행시간은 최선의 경우 `O(n)`, 최악의 경우 `O(n^2)`, 평균의 경우 `O(n^2)`이다.

### 버블 정렬 특징

1. 선택 정렬에 비해 원소의 교환이 많이 발생함: 수행시간이 최악인 경우 (`O(n^2)`) 선택 정렬보다 비효율적임
2. 안정적인 정렬임
3. 제자리 정렬임