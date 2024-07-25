import os
import subprocess
import tempfile

TOM_TOM_API_KEY_ENV_KEY = "TOM_TOM_API_KEY"

def run_docker_compose(secret):
    # Create a temporary file to store the secret
    with tempfile.NamedTemporaryFile(mode='w+', delete=False) as temp_file:
        temp_file.write(secret)
        temp_file_path = temp_file.name

    try:
        # Set the environment variable with the path to the secret file
        os.environ['SECRET_FILE_PATH'] = temp_file_path

        # Run docker-compose up
        subprocess.run(['docker-compose', 'up', '--build'], check=True)
    except subprocess.CalledProcessError as e:
        print(f"Error running Docker Compose: {e}")
    except FileNotFoundError:
        print("Docker Compose not found. Make sure it's installed and in your PATH.")
    finally:
        # Clean up: remove the temporary file
        os.unlink(temp_file_path)

def main():
    key = os.getenv(TOM_TOM_API_KEY_ENV_KEY)
    if not key:
        print(f"{TOM_TOM_API_KEY_ENV_KEY} environment variable not set.")
        return

    # Run Docker Compose with the secret
    run_docker_compose(key)

if __name__ == "__main__":
    main()