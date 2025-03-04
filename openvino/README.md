# Installing Openvino on my Intel AI PC

## Table of contents

- [Installing Openvino on my Intel AI PC](#installing-openvino-on-my-intel-ai-pc)
  - [Table of contents](#table-of-contents)
  - [1. Install Python and Git](#1-install-python-and-git)
  - [2. Install Drivers for GPU, and NPU (AI PC)](#2-install-drivers-for-gpu-and-npu-ai-pc)
  - [3. Install C++ Redistributable (required), FFMPEG (optional)](#3-install-c-redistributable-required-ffmpeg-optional)
  - [4. Install optimum-intel](#4-install-optimum-intel)
  - [5. Clone the Repository](#5-clone-the-repository)
  - [6. Install the Packages](#6-install-the-packages)
  - [7. Launch Notebooks](#7-launch-notebooks)
  - [8. To open the notebooks](#8-to-open-the-notebooks)



For the notebooks installation I followed [this guide](https://github.com/openvinotoolkit/openvino_notebooks/wiki/Windows)

## 1. Install Python and Git
1.1 Install Python
NOTE: ⚠️The version of Python that is available in the Microsoft Store is not recommended.⚠️ So I had to uninstall the latest version that I had and install [Python 3.10](https://www.python.org/ftp/python/3.10.11/python-3.10.11-amd64.exe)
Don't forget to:
- Check the box to add Python to your PATH, and to install py. 
- At the end of the installer, there is an option to disable the PATH length limit. It is recommended to click this.
1.2 Install Git
Download GIT from this link https://github.com/git-for-windows/git/releases/download/v2.45.1.windows.1/Git-2.45.1-64-bit.exe

## 2. Install Drivers for GPU, and NPU (AI PC)
I did a "Clean Install" of the WHQL Certified GPU driver to ensure the underlying libraries are correctly configured. https://www.intel.com/content/www/us/en/download/785597/834050/intel-arc-iris-xe-graphics-windows.html

Additionally, for AI PC users, I installed the latest NPU driver (or last known working Intel® NPU Driver - Windows* 32.0.100.3053) to avoid any potential issues in compiling NPU kernels. https://www.intel.com/content/www/us/en/download/794734/intel-npu-driver-windows.html
The guide on how to install it, is https://www.intel.com/content/www/us/en/support/articles/000099083/processors/intel-core-ultra-processors.html

## 3. Install C++ Redistributable (required), FFMPEG (optional)
- Downloaded [Microsoft Visual C++ Redistributable](https://aka.ms/vs/17/release/vc_redist.x64.exe).
- Downloaded FFMPEG from https://www.gyan.dev/ffmpeg/builds/
    - Clicked on "ffmpeg-release-full.zip" under the "Release Builds" section.
    - Extracted the ZIP file in the location C:\ffmpeg.
    - Added FFmpeg to System PATH 
      - Opened File Explorer and navigate to C:\ffmpeg\bin.
      - Copied the path (C:\ffmpeg\bin).
      - Opened System Environment Variables
      - Press Win + R, type sysdm.cpl, and hit Enter.
      - Go to the Advanced tab → Click Environment Variables.
      - Under System Variables, find Path, select it, and click Edit.
      - Click New, paste C:\ffmpeg\bin, and press OK.
      - Close and restart Command Prompt.
      - Verify installation with `ffmpeg -version`


## 4. Install optimum-intel
Then, in command prompt, created a python env and activate it
```bash
python -m venv optimum_intel
openvino_env\Scripts\activate
```

 and install the following
https://github.com/huggingface/optimum-intel
```bash
pip install --upgrade --upgrade-strategy eager "optimum[neural-compressor]"
pip install --upgrade --upgrade-strategy eager "optimum[openvino]"
pip install --upgrade --upgrade-strategy eager "optimum[ipex]"
```

## 5. Clone the Repository

```bash
git clone --depth=1 https://github.com/openvinotoolkit/openvino_notebooks.git
cd openvino_notebooks
```

## 6. Install the Packages

```bash
python -m pip install --upgrade pip wheel setuptools
pip install -r requirements.txt
```

## 7. Launch Notebooks
To launch a single notebook, like the PyTorch to OpenVINO notebook
```bash
jupyter lab notebooks/pytorch-to-openvino/pytorch-to-openvino.ipynb
```
To launch all notebooks in Jupyter Lab (localhost only)
```bash
jupyter lab notebooks
```
To launch all notebooks available from any host
```bash
jupyter lab notebooks --ip 0.0.0.0
```
In Jupyter Lab, select a notebook from the file browser using the left sidebar. Each notebook is located in a subdirectory within the notebooks directory.


## 8. To open the notebooks
```bash
optimum_intel\Scripts\activate
cd openvino_notebooks
jupyter lab notebooks
```

My goal is to use the [llm-question-answering notebook](https://github.com/openvinotoolkit/openvino_notebooks/tree/latest/notebooks/llm-question-answering) with mistral-7b.


