# 2026-02-20

Kopia `python-masterclass` -> Read-Only here.

# AIDUMP

Przydatna komenda:

```sh
vi ~/.bashrc
```

```sh
alias aidump='{ \
  echo "<<<PROJECT_TREE>>>"; \
  tree -a -I ".git|__pycache__|node_modules|.venv|venv|pgdata"; \
  echo "<<<END_PROJECT_TREE>>>"; \
  find . -type f \
    ! -path "./.git/*" \
    ! -path "*/__pycache__/*" \
    ! -path "*/node_modules/*" \
    ! -path "*/.venv/*" \
    ! -path "*/venv/*" \
    ! -path "*/pgdata/*" \
    ! -iname "*.png" \
    ! -iname "*.jpg" \
    ! -iname "*.jpeg" \
    ! -iname "*.gif" \
    ! -iname "*.webp" \
    ! -iname "*.svg" \
    ! -iname "*.ico" \
    ! -iname "*.pdf" \
    ! -iname "*.zip" \
    ! -iname "*.tar" \
    ! -iname "*.gz" \
    ! -iname "*.7z" \
    ! -iname "*.mp4" \
    ! -iname "*.mp3" \
    | sort \
    | while read -r f; do \
        echo; \
        echo "<<<FILE:$f>>>"; \
        sed -n "1,300p" "$f"; \
        echo "<<<END_FILE:$f>>>"; \
      done; \
}'
```