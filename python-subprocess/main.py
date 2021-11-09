import subprocess

result = subprocess.run(
    ["git", "ls-files"],
    capture_output=True,
    text=True
)

result.check_returncode()
print(result.stdout)
