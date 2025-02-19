import openvino as ov
import numpy as np

# Initialize OpenVINO Core
core = ov.Core()
print(f"Available devices: {core.available_devices}")

# Test GPU compute capability
try:
    # Create a simple test tensor
    data = np.array([[1, 2, 3]], dtype=np.float32)
    
    # Create a simple model using OpenVINO operations
    input_node = ov.opset8.parameter([1, 3], dtype=np.float32)
    const_node = ov.opset8.constant(2, dtype=np.float32)
    mul_node = ov.opset8.multiply(input_node, const_node)
    model = ov.Model(mul_node, [input_node], "test_model")
    
    # Compile for GPU
    compiled_model = core.compile_model(model, "GPU")
    
    # Create inference request
    infer_request = compiled_model.create_infer_request()
    
    # Run inference
    infer_request.set_input_tensor(ov.Tensor(data))
    infer_request.infer()
    
    # Get result
    output = infer_request.get_output_tensor().data
    print(f"Test computation result: {output}")
    print("GPU compute test successful!")
except Exception as e:
    print(f"Error during GPU test: {e}") 