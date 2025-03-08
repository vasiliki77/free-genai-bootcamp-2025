# Use an official lightweight Python image
FROM python:3.11-slim

# Set the working directory in the container
WORKDIR /app

# Copy the app directory contents into the container
COPY app/ .

# Install system dependencies
RUN apt-get update && apt-get install -y \
    git \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose Streamlit port
EXPOSE 8502

# Command to run the app
CMD ["streamlit", "run", "diacritics.py", "--server.port=8502", "--server.address=0.0.0.0"]