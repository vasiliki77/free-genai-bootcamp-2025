FROM python:3.10-slim

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential curl net-tools \
    && rm -rf /var/lib/apt/lists/*

# Copy necessary files (directly from current directory)
COPY requirements.txt /app/
COPY *.py /app/

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose the port
EXPOSE 8080

# Healthcheck to verify the service is operational
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

# Command to run the application
CMD ["uvicorn", "opea_translation_microservice:app", "--host", "0.0.0.0", "--port", "8080"]