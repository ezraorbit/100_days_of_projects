import math

height = float(input("enter your height in m: "))
weight = float(input("enter your weight in kg: "))

bmi = round(weight / math.pow(height, 2), 2)

if(bmi < 18.5):
    print(f"You have a bmi of {bmi}, you are underweight")
elif(bmi < 25):
    print(f"You have a bmi of {bmi}, you are a normal weight")
elif(bmi < 30):
    print(f"You have a bmi of {bmi}, you are overweight")
elif(bmi < 35):
    print(f"You have a bmi of {bmi}, you are obese")
else:
    print(f"You have a bmi of {bmi}, you are clinically obese")