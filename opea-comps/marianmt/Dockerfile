FROM ubuntu:22.04

# Install system dependencies
RUN apt-get update && apt-get install -y \
    python3-pip \
    python3-dev \
    git

WORKDIR /app

# Install Python dependencies
COPY requirements.txt .
RUN pip3 install -r requirements.txt

# Copy service files
COPY src/ .

# Expose the port
EXPOSE 8000

# Run the service
CMD ["python3", "app.py"] 