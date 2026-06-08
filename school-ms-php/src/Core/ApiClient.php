<?php

namespace Core;

class ApiClient
{
    private string $baseUrl;
    private ?string $token;

    public function __construct(?string $token = null)
    {
        $this->baseUrl = Config::apiUrl();
        $this->token   = $token ?? Session::get('access_token');
    }

    public function get(string $path, array $query = []): array
    {
        $url = $this->baseUrl . $path;
        if ($query) $url .= '?' . http_build_query($query);
        return $this->request('GET', $url);
    }

    public function post(string $path, array $data = []): array
    {
        return $this->request('POST', $this->baseUrl . $path, $data);
    }

    public function put(string $path, array $data = []): array
    {
        return $this->request('PUT', $this->baseUrl . $path, $data);
    }

    public function patch(string $path, array $data = []): array
    {
        return $this->request('PATCH', $this->baseUrl . $path, $data);
    }

    public function delete(string $path): array
    {
        return $this->request('DELETE', $this->baseUrl . $path);
    }

    private function request(string $method, string $url, array $data = []): array
    {
        $headers = [
            'Content-Type: application/json',
            'Accept: application/json',
            // Forward the browser's domain as Host so the Go middleware
            // can resolve the tenant. Without this, cURL sends "localhost:8080"
            // and the tenant lookup finds nothing.
            'Host: ' . Config::currentDomain(),
        ];

        if ($this->token) {
            $headers[] = 'Authorization: Bearer ' . $this->token;
        }

        $ch = curl_init($url);
        curl_setopt_array($ch, [
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_HTTPHEADER     => $headers,
            CURLOPT_CUSTOMREQUEST  => $method,
            CURLOPT_TIMEOUT        => 15,
        ]);

        if (in_array($method, ['POST', 'PUT', 'PATCH']) && $data) {
            curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
        }

        $response = curl_exec($ch);
        $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        $error    = curl_error($ch);
        curl_close($ch);

        if ($error) {
            return ['success' => false, 'error' => 'API connection failed: ' . $error];
        }

        $decoded = json_decode($response, true);

        if ($httpCode === 401) {
            Session::destroy();
            header('Location: /auth/login');
            exit;
        }

        return $decoded ?? ['success' => false, 'error' => 'Invalid API response (HTTP ' . $httpCode . ')'];
    }
}