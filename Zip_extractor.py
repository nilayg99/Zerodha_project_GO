import os
import zipfile
import datetime

# Get the current date and the previous date in DDMMYY format
current_date = datetime.datetime.now()
previous_date = current_date - datetime.timedelta(days=1)
previous_date_str = previous_date.strftime('%d%m%y')
print("Previous date:", previous_date_str)

# Define the path to the ZIP file
zip_file_path = 'D:/Project/Project_New/Zerodha_project/Zerodha_project_GO/BSE_File_Saved/'
zip_file = os.path.join(zip_file_path, f"EQ{previous_date_str}_CSV.ZIP")
print("ZIP file path:", zip_file)

# Define the path to the extracted files
extracted_files_path = 'D:/Project/Project_New/Zerodha_project/Zerodha_project_GO/BSE_File_Saved'
print("Extracted files path:", extracted_files_path)

# Define the CSV file name format
csv_file_name = f'EQ{previous_date_str}.csv'

# Define the full CSV file path
csv_file_path = os.path.join(extracted_files_path, csv_file_name)
print("CSV file path:", csv_file_path)

# Check if the CSV file exists
if os.path.exists(csv_file_path):
    print(f"The CSV file '{csv_file_path}' is present.")
else:
    # Extract the ZIP file if the CSV file is not present
    try:
        with zipfile.ZipFile(zip_file, 'r') as zip_ref:
            zip_ref.extractall(extracted_files_path)
        print(f"The CSV file '{csv_file_path}' was extracted from the ZIP file '{zip_file}'.")
    except zipfile.BadZipFile:
        print(f"Error: The ZIP file '{zip_file}' is corrupted or not a valid ZIP file.")
    except PermissionError:
        print(f"Error: You don't have permission to extract files to '{extracted_files_path}'.")
    except Exception as e:
        print(f"Error: {e}")
