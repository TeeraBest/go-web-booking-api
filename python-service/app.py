from flask import Flask, jsonify
import pandas as pd # data processing, CSV file I/O (e.g. pd.read_csv)
import joblib
from flask import request

app = Flask(__name__)

@app.route('/hello')
def hello():
    msg = "Hello from Python!"
    return jsonify(message=msg)

@app.route('/predict', methods=['POST'])
def predictDataFromParams():
    """ Example of a cURL command to test this endpoint:
    curl -X POST http://localhost:5001/predict \
    -H "Content-Type: application/json" \
    -d '{
        "Time": 6986,
        "V1": -4.39797444171999,
        "V2": 1.35836702839758,
        "V3": -2.5928442182573,
        "V4": 2.67978696694832,
        "V5": -1.12813094208956,
        "V6": -1.70653638774951,
        "V7": -3.49619729302467,
        "V8": -0.248777743025673,
        "V9": -0.24776789948008,
        "V10": -4.80163740602813,
        "V11": 4.89584422347523,
        "V12": -10.9128193194019,
        "V13": 0.184371685834387,
        "V14": -6.77109672468083,
        "V15": -0.00732618257771211,
        "V16": -7.35808322132346,
        "V17": -12.5984185405511,
        "V18": -5.13154862842983,
        "V19": 0.308333945758691,
        "V20": -0.17160787864796,
        "V21": 0.573574068424352,
        "V22": 0.176967718048195,
        "V23": -0.436206883597401,
        "V24": -0.0535018648884285,
        "V25": 0.252405261951833,
        "V26": -0.657487754764504,
        "V27": -0.827135714578603,
        "V28": 0.849573379985768,
        "Amount": 59
    }'
    """
    """
    Predict data from JSON parameters.
    Expects a JSON object in the request body with the input data for prediction.
    Returns a JSON response with the prediction result.
    """
    data = request.get_json()
    if not data:
        return jsonify(error="No input data provided"), 400
    try:
        prediction = preditModel(data)
    except Exception as e:
        return jsonify(error=str(e)), 400
    return jsonify(prediction=str(prediction))

def preditModel(data):
    """
    Predict using the loaded model and provided data dictionary.
    :param data: dict with feature names as keys.
    :return: predicted class/value.
    """
    # Required feature names
    sample_features = [
        'Time', 'V1', 'V2', 'V3', 'V4', 'V5', 'V6', 'V7', 'V8', 'V9',
        'V10', 'V11', 'V12', 'V13', 'V14', 'V15', 'V16', 'V17', 'V18', 'V19',
        'V20', 'V21', 'V22', 'V23', 'V24', 'V25', 'V26', 'V27', 'V28', 'Amount'
    ]
    if not isinstance(data, dict):
        raise TypeError("Input data must be a dictionary with feature names as keys.")
    if not all(feature in data for feature in sample_features):
        raise ValueError(f"Data must contain the following features: {', '.join(sample_features)}")
    # Load model
    loaded_model = joblib.load('MLmodels/model.joblib')
    # Prepare DataFrame
    
    sample_data = pd.DataFrame([data], columns=sample_features)
    # Predict
    sample_pred = loaded_model.predict(sample_data)
    return sample_pred[0]

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
