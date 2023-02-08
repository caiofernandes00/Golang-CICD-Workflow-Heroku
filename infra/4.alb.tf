resource "aws_security_group" "golang_app_alb" {
  name        = "${aws_ecr_repository.my_app.name}-alb-sg"
  description = "Security group for the Golang API ALB"
  vpc_id      = aws_vpc.golang_app.id

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_alb" "golang_app" {
  name            = "${aws_ecr_repository.my_app.name}-alb"
  internal        = false
  security_groups = [aws_security_group.golang_app_alb.id]

  subnets = aws_subnet.golang_app.*.id
  tags = {
    Name = aws_ecr_repository.my_app.name
  }
}

resource "aws_alb_target_group" "golang_app" {
  name     = "${aws_ecr_repository.my_app.name}-target-group"
  port     = 8080
  protocol = "HTTP"
  vpc_id   = aws_vpc.golang_app.id

  health_check {
    path                = "/health"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 3
    unhealthy_threshold = 3
    matcher             = "200-299"
  }

  depends_on = [aws_alb.golang_app]

  tags = {
    Name = aws_ecr_repository.my_app.name
  }
}

resource "aws_alb_listener" "golang_app" {
  load_balancer_arn = aws_alb.golang_app.arn
  port              = "8080"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.golang_app.arn
  }
}
