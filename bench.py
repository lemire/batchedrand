# uv run --with=pandas --with=matplotlib --with=tabulate bench.py
import subprocess
import sys
import re
import pandas as pd
import matplotlib.pyplot as plt

plt.rcParams.update({'font.size': 14})

# Run go test -bench=. and capture output
process = subprocess.Popen(
    ['go', 'test', '-bench=.'],
    stdout=subprocess.PIPE,
    stderr=subprocess.STDOUT,
    text=True,
    bufsize=1  # Line buffered
)

# Collect full output for later parsing
full_output = []
print("Running benchmarks...\n")

# Print each line in real-time
for line in process.stdout:
    print(line, end='')  # Already includes newline
    full_output.append(line)

process.wait()

if process.returncode != 0:
    print(f"\nBenchmark command failed with return code {process.returncode}")
    sys.exit(1)

print("\nBenchmarks completed. Parsing results...\n")

# Parse the collected output
data = []
pattern = re.compile(r'Benchmark(\w+)/(\w+)_size_(\d+)-\d+\s+\d+\s+([\d\.]+) ns/op')

for line in full_output:
    match = pattern.match(line.strip())
    if match:
        bench_base = match.group(1)
        method = match.group(2)
        size = int(match.group(3))
        ns_op = float(match.group(4))
        
        bench_type = bench_base + 'Shuffle' if bench_base in ['ChaCha', 'PCG'] else bench_base
        
        data.append({
            'Benchmark': bench_type,
            'Method': method,
            'Size': size,
            'ns/op': ns_op
        })

if not data:
    print("No benchmark data found in output.")
    sys.exit(1)

df = pd.DataFrame(data)
df = df.sort_values(['Benchmark', 'Size', 'Method'])

# Pivot to get Batched and Standard in columns
pivot_df = df.pivot(index=['Benchmark', 'Size'], columns='Method', values='ns/op').reset_index()

# Compute ns/item and speedup
pivot_df['Batched (ns/item)'] = pivot_df['Batched'] / pivot_df['Size']
pivot_df['Standard (ns/item)'] = pivot_df['Standard'] / pivot_df['Size']
pivot_df['speedup'] = pivot_df['Standard'] / pivot_df['Batched']

# Format to ensure one decimal place
pivot_df['Batched (ns/item)'] = pivot_df['Batched (ns/item)'].map(lambda x: f"{x:.1f}")
pivot_df['Standard (ns/item)'] = pivot_df['Standard (ns/item)'].map(lambda x: f"{x:.1f}")
pivot_df['speedup'] = pivot_df['speedup'].map(lambda x: f"{x:.1f}")

# Table with desired columns
table_df = pivot_df[['Benchmark', 'Size', 'Batched (ns/item)', 'Standard (ns/item)', 'speedup']]

# Print markdown table
print("### Benchmark results table")
print(table_df.to_markdown(index=False))
print("\n")

# Generate and save bar charts
benchmark_types = pivot_df['Benchmark'].unique()
sizes = sorted(pivot_df['Size'].unique())

for bench_type in benchmark_types:
    plt.figure(figsize=(10, 6))
    subset = pivot_df[pivot_df['Benchmark'] == bench_type].sort_values('Size')
    
    x = range(len(sizes))
    width = 0.35
    
    plt.bar([i - width/2 for i in x], subset['Batched (ns/item)'], width, label='Batched', color='darkblue')
    plt.bar([i + width/2 for i in x], subset['Standard (ns/item)'], width, label='Standard', color='red')
    
    plt.xlabel('Array size')
    plt.ylabel('Time (ns/item)')
    plt.title(f'{bench_type} benchmark: Batched vs Standard')
    plt.xticks(x, sizes)
    plt.legend(frameon=False)
    plt.grid(True, which="both", ls="--", alpha=0.7)
    
    # Remove top and right spines
    ax = plt.gca()
    ax.spines['top'].set_visible(False)
    ax.spines['right'].set_visible(False)
    
    plt.tight_layout()
    
    filename = f"{bench_type.lower()}_benchmark.png"
    plt.savefig(filename)
    print(f"Saved chart: {filename}")
    plt.close()

print("\nAll done.")