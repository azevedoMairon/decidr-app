if [ -z "$BACKEND_URL" ]; then
  echo "Erro: BACKEND_URL não está definido."
  exit 1
fi

envsubst '$BACKEND_URL' < /etc/nginx/templates/nginx.conf.template > /etc/nginx/conf.d/default.conf

exec nginx -g 'daemon off;'