# 탐색 (Search)

탐색이란 여러 원소로 구성된 데이터에서 원하는 원소를 찾는 것이다.

탐색은 저장 형태에 따라 구분된다. 저장형태는 list, tree, graph, table, string 등이 있다.

## 순차 탐색

## 이진 탐색

이진 탐색은 정렬된 리스트 형태로 주어진 원소들을 절반씩 나누어 가운데 원소를 기준으로 원하는 키값을 갖는 원소를 찾는 탐색방식이다.

`{ 7, 15, 22, 30, 35, 40, 44, 55, 88 }`의 배열에서 44를 찾는다고 하면

1. 먼저 가운데 원소인 35와 44를 비교한다.
2. 35가 44보다 작으므로 35의 오른쪽 부분에 대해 탐색을 진행한다.
3. `{ 40, 44, 55, 88 }`의 배열을 대상으로 가운데 원소인 55와 44를 비교한다.
4. 55가 44보다 크므로 55의 왼쪽 부분에 대해 탐색을 진행한다.
5. `{ 40, 44 }`의 배열을 대상으로 가운데 원소인 44를 찾을 수 있다.

### 이진 탐색 알고리즘

```python
def binary_search(dlist, target):
    left = 0
    right = len(dlist) - 1
    while (left <= right):
        mid = (left + right) // 2
        if (dlist[mid] == target):
            return mid
        elif (dlist[mid] < target):
            left = mid + 1
        else:
            right = mid - 1
    return None
```

순환(재귀) 알고리즘으로 구현할 수 있다.

```python
def binary_search_recursion(dlist, target, left, right):
    if (left > right): return None
    mid = (left + right) // 2

    if (dlist[mid] == target):
        return mid
    elif (dlist[mid] < target):
        left = mid + 1
    else:
        right = mid - 1

    return binary_search_recursion(dlist, target, left, right)
```

### 이진 탐색 성능 분석

이진 탐색 알고리즘은 하나의 반복문으로 구성된다.
반복마다 원소의 개수가 절반으로 감소한다. 따라서 최악, 평균 수행시간 모두 `O(logn)`이다.

### 이진 탐색 특징

이진 탐색은 정렬된 리스트가 입력값으로 들어와야 적용 가능하다
