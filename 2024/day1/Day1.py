# read file data into 2 lists.

def read_columns_into_lists(filename):
    left_list = []
    right_list = []
    with open(filename, 'r') as file:
        for line in file:
            # Split the line into two parts based on whitespace
            left, right = map(int, line.split())
            left_list.append(left)
            right_list.append(right)
    return left_list, right_list

# Sort those lists smallest to largest.
def sort_list(list):
  return list.sort()

# calculate and sum differences.
def calculate_dif(left, right):
  total = 0
  
  for index, val in enumerate(left):
    diff = val - right[index]
    if diff > 0:
      total += diff
    else:
      total += 0 - diff
      
  return total

def create_dictionary(list):
  dict = {}
  
  for item in list:
    if dict.get(item):
      dict[item] += 1
    else:
      dict[item] = 1
  
  return dict

def calculate_similarity(dict, list):
  similarity = 0
  
  for item in list:
    count_or_something = dict.get(item)
    if count_or_something != None:
      similarity += item * count_or_something
      
  return similarity

filename = 'data'
left, right = read_columns_into_lists(filename)
dict = create_dictionary(right)
print(calculate_similarity(dict, left))