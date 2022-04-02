from flask import Flask
from dotenv import load_dotenv

app = Flask(__name__)

load_dotenv()
  
@app.route('/')
def hello():
    return "Welcome to my SCA Cloud School Application, this is my first assessment."
    