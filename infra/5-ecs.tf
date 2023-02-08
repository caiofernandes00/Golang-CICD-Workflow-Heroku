resource "aws_ecs_cluster" "golang_app" {
  name = "${aws_ecr_repository.my_app.name}-cluster"
}

resource "aws_ecs_task_definition" "golang_app" {
  family = aws_ecr_repository.my_app.name
  container_definitions = jsonencode([
    {
      name      = aws_ecr_repository.my_app.name
      image     = "${aws_ecr_repository.my_app.repository_url}:latest"
      cpu       = 128
      memory    = 128
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
          protocol      = "tcp"
        }
      ]
      environment = [
        {
          name  = "PORT"
          value = "8080"
        }
      ]
    }
  ])
}

resource "aws_ecs_service" "golang_app" {
  name            = "${aws_ecr_repository.my_app.name}-service"
  cluster         = aws_ecs_cluster.golang_app.id
  task_definition = aws_ecs_task_definition.golang_app.arn
  desired_count   = 2
  deployment_controller {
    type = "ECS"
  }
  load_balancer {
    target_group_arn = aws_alb_target_group.golang_app.arn

    container_name = aws_ecr_repository.my_app.name
    container_port = 8080
  }
}
