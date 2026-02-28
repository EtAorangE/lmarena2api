#!/usr/bin/env python3
"""
自动拉取 lmarena 模型列表并更新 constants.go

使用方法:
    python3 fetch_models.py --cookie "your-cookie" --cf-clearance "your-cf-clearance"

或者设置环境变量:
    export LA_COOKIE="your-cookie"
    export CF_CLEARANCE="your-cf-clearance"
    python3 fetch_models.py
"""

import argparse
import json
import os
import re
import requests
from typing import Dict, List

LMARENA_BASE_URL = "https://canary.lmarena.ai"

def fetch_models_from_page(cookie: str, cf_clearance: str) -> List[Dict]:
    """从 lmarena 页面获取模型列表"""
    headers = {
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36",
        "Accept": "text/html,application/xhtml+xml",
        "Cookie": f"cf_clearance={cf_clearance}; arena-auth-prod-v1={cookie}",
    }
    
    try:
        response = requests.get(f"{LMARENA_BASE_URL}/", headers=headers, timeout=30)
        response.raise_for_status()
        
        models = []
        model_pattern = r'"modelId"\s*:\s*"([^"]+)"'
        model_ids = re.findall(model_pattern, response.text)
        
        for model_id in set(model_ids):
            models.append({"id": model_id, "name": model_id})
        
        return models
    except Exception as e:
        print(f"获取模型列表失败: {e}")
        return []

def update_constants_go(models: List[Dict], output_path: str):
    """更新 constants.go 文件"""
    with open(output_path, 'r') as f:
        content = f.read()
    
    model_entries = []
    for model in models:
        name = model.get("name", model.get("id", ""))
        model_id = model.get("id", "")
        model_type = model.get("type", "chat")
        
        if name and model_id:
            model_entries.append(f'\t"{name}": {{"{name}", "{model_id}", "{model_type}"}},')
    
    pattern = r'var ModelRegistry = map\[string\]ModelInfo\{[^}]+\}'
    replacement = 'var ModelRegistry = map[string]ModelInfo{\n' + '\n'.join(model_entries) + '\n}'
    
    new_content = re.sub(pattern, replacement, content, flags=re.DOTALL)
    
    with open(output_path, 'w') as f:
        f.write(new_content)
    
    print(f"✅ 已更新 {output_path}")

def main():
    parser = argparse.ArgumentParser(description="拉取 lmarena 模型列表")
    parser.add_argument("--cookie", default=os.environ.get("LA_COOKIE", ""), help="LA Cookie")
    parser.add_argument("--cf-clearance", default=os.environ.get("CF_CLEARANCE", ""), help="CF Clearance")
    parser.add_argument("--output", default="common/constants.go", help="输出文件路径")
    parser.add_argument("--dry-run", action="store_true", help="只显示结果，不更新文件")
    
    args = parser.parse_args()
    
    if not args.cookie:
        print("错误: 请提供 LA_COOKIE")
        return
    
    print("正在获取模型列表...")
    
    models = fetch_models_from_page(args.cookie, args.cf_clearance)
    
    if not models:
        print("❌ 无法获取模型列表")
        return
    
    print(f"找到 {len(models)} 个模型")
    
    if args.dry_run:
        print("\n模型列表:")
        for model in models:
            print(f"  - {model.get('name', model.get('id'))}")
    else:
        update_constants_go(models, args.output)

if __name__ == "__main__":
    main()
