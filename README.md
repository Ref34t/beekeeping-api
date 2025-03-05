# Beekeeping API

A simple RESTful API for retrieving senseBox data related to beekeeping chores. Data is sourced from openSenseMap and returned as JSON.

## Tools
- **Git**: Version control.
- **VS Code**: IDE for development.
- **Docker**: Containerization.

## Endpoints

### GET /api/v1/sensebox/{id}
Retrieve data for a specific senseBox by its `id`.

#### Example Response:
```json
{
  "id": "5eba5fbad46fb8001b799786",
  "temperature": 22.5,
  "humidity": 60.0
}