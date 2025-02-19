# Copyright (C) 2024 Intel Corporation
# SPDX-License-Identifier: Apache-2.0

# Stage 1: base setup
FROM python:3.11-slim AS base

# Get security updates
RUN apt-get update && apt-get upgrade -y && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

ENV HOME=/home/user
RUN useradd -m -s /bin/bash user && \
    mkdir -p $HOME && \
    chown -R user $HOME

WORKDIR $HOME

# Stage 2: Install OPEA components
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Stage 3: Add application code
COPY app.py .

USER user
CMD ["python", "app.py"] 