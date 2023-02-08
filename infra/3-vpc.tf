resource "aws_vpc" "golang_app" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = aws_ecr_repository.my_app.name
  }
}

resource "aws_internet_gateway" "golang_app_igw" {
  vpc_id = aws_vpc.golang_app.id

}

resource "aws_subnet" "golang_app" {
  count             = 2
  vpc_id            = aws_vpc.golang_app.id
  cidr_block        = "10.0.${count.index + 1}.0/24"
  availability_zone = var.availability_zones[count.index]
  tags = {
    Name = "${aws_ecr_repository.my_app.name}-subnet-${count.index}"
  }
}

