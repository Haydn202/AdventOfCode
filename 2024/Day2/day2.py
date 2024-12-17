def read_in_reports(filename):
  reports = []
  
  with open(filename, 'r') as file:
    for line in file:
      reports.append(list(map(int, line.split())))
      
  return reports

def check_reports(reports):
  total_safe = 0
  for report in reports:
    if check_report(report) is True:
      total_safe += 1
  
  return total_safe
  
def check_report(report):
  failed_count = 0
  
  dict = {
    True: 0,
    False: 0
  }
  
  prev = report[0]
  
  is_first = True
  
  for item in report[1:]:
    if failed_count > 1:
      print("failed")
      return False
    
    difference = item - prev
    
    increasing = difference > 0
    
    absolute = abs(prev - item)
    
    to_big = absolute > 3 or absolute < 1
    
    if to_big and is_first:
      prev = item
      
    is_first = False
    if difference != 0 and not to_big:
      dict[increasing] += 1
    
    if absolute > 3 or absolute < 1:
      failed_count += 1
      print("Abs prev: ", prev, "Current: ", item)
      print("abs error: ", failed_count)
      continue

    prev = item
    
  if failed_count > 1:
    print("failed")
    return False
  print("dict: ", dict)
  print("valid direction: ", is_valid_direction(dict, failed_count))
  return is_valid_direction(dict, failed_count)

# arr int
# if all same direction all good

def is_valid_direction(dict, failed_count):
  print("min val: ", min(dict.values()))
  if min(dict.values()) > 0 and failed_count != 0:
    return False
  
  if min(dict.values()) > 1:
    print("dict: ", dict)
    return False
  
  return True

def check_direction(increasing, difference):
  if increasing and difference > 0:
    return True
  
  if increasing is False and difference < 0:
    return True
    
  return False
  
reports = read_in_reports("data")
safe_reports = check_reports(reports)
print(safe_reports)