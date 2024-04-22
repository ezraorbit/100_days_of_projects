# Tip calculator
print("Welcome to the tip calculator")
total_bill = float(input("What was the total bill? $"))
tip_percentage = float(input("What percentage tip would you like to give? 10, 12, or 15? ")) / 100.0
total_people = float(input("How many people to split the bill? "))

tip_total = total_bill * tip_percentage
total_with_tip = total_bill + tip_total
total_to_pay = total_with_tip / total_people

print("Each person should pay: %.2f" % total_to_pay)

