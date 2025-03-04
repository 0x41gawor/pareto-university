from googletrans import Translator


translator = Translator()

f = open("temp.txt", "r")
for x in f:
  print(x)
  xt = translator.translate(x, dest='es', src='en')
  print(f'{x}: {xt}')
