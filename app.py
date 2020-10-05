#!flask/bin/python
from flask import Flask, jsonify
from flask import abort
from flask import make_response
from flask import request
from flask_cors import CORS
from ctypes import *

app = Flask(__name__)
CORS(app)

lib = cdll.LoadLibrary("./calculations.so")

class GoSlice(Structure):
    _fields_ = [("data", POINTER(c_double)), ("len", c_longlong), ("cap", c_longlong)]

class MeanStddev_return(Structure):
    _fields_ = [("r0", c_double), ("r1", c_double)]

class Quartile_return(Structure):
    _fields_ = [("r0", c_double), ("r1", c_double), ("r2", c_double)]

class Mode_return(Structure):
    _fields_ = [("r0", c_double), ("r1", c_longlong)]

lib.MeanStddev.argtypes = [GoSlice, c_longlong]
lib.MeanStddev.restype = MeanStddev_return
lib.QuickSort.argtypes = [GoSlice]
lib.QuickSort.restype = None
lib.Quartile.argtypes = [GoSlice]
lib.Quartile.restype = Quartile_return
lib.Mode.argtypes = [GoSlice]
lib.Mode.restype = Mode_return

tasks = [
    {
        'id': 1,
        'arr': [0, 0, 0, 0, 0, 0],
        'data': []
    }
]

@app.route('/stats/calc', methods=['GET'])
def get_tasks():
    return jsonify({'tasks': tasks})

@app.route('/stats/calc/<int:task_id>', methods=['GET'])
def get_task(task_id):
    task = [task for task in tasks if task['id'] == task_id]
    if len(task) == 0:
        abort(404)
    return jsonify({'task': task[0]})

@app.errorhandler(404)
def not_found(error):
    return make_response(jsonify({'error': 'Not found'}), 404)

@app.route('/stats/calc', methods=['POST'])
def create_task():
    print(request.json)
    if not request.json or not 'title' in request.json:
        abort(400)
    task = {
        'id': tasks[-1]['id'] + 1,
        'arr': request.json['arr']
    }
    tasks.append(task)
    return jsonify({'task': task}), 201

@app.route('/stats/calc/<int:task_id>', methods=['PUT'])
def update_task(task_id):
    task = [task for task in tasks if task['id'] == task_id]
    if len(task) == 0:
        abort(404)
    if not request.json:
        abort(400)
    # if 'arr' in request.json:
    #     abort(400)
    task[0]['arr'] = request.json.get('arr', task[0]['arr'])
    task[0]['data'] = get_summary(task[0]['arr'])
    return jsonify(task[0])

@app.route('/stats/calc/<int:task_id>', methods=['DELETE'])
def delete_task(task_id):
    task = [task for task in tasks if task['id'] == task_id]
    if len(task) == 0:
        abort(404)
    tasks.remove(task[0])
    return jsonify({'result': True})

def get_summary(arr):
    nums = GoSlice((c_double * len(arr))(*arr), len(arr), len(arr))
    std_dev = lib.MeanStddev(nums, 1)
    mean = std_dev.r0
    std = std_dev.r1

    quartile = lib.Quartile(nums)
    q1 = quartile.r0
    q2 = quartile.r1
    q3 = quartile.r2

    mode = lib.Mode(nums)
    freq = mode.r0
    times = mode.r1
    
    result = [mean, std]
    result.append(nums.data[0])
    result.append(q1)
    result.append(q2)
    result.append(q3)
    result.append(nums.data[nums.len-1])
    if times <= 1:
        result.append(None)
    else:
        result.append(freq)
        result.append(times)

    for i in range(nums.len):
        arr[i] = nums.data[i]
    return result


if __name__ == '__main__':
    app.run(debug=False)